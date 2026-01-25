package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/mailer"
	"github.com/cconner57/adoption-os/backend/internal/notifier"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	assetsDir string
}

type application struct {
	config   config
	logger   *slog.Logger
	models   data.Models
	mailer   mailer.Mailer
	notifier *notifier.Notifier
	db       *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Note: .env file not found (relying on system env vars)")
	}

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	flag.StringVar(&cfg.smtp.host, "smtp-host", os.Getenv("SMTP_SERVER"), "SMTP host")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	flag.IntVar(&cfg.smtp.port, "smtp-port", port, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv("SMTP_LOGIN"), "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv("SMTP_KEY"), "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", os.Getenv("SMTP_SENDER"), "SMTP sender email")

	flag.StringVar(&cfg.assetsDir, "assets-dir", os.Getenv("ASSETS_DIR"), "Directory for storing uploaded assets")
	if cfg.assetsDir == "" {
		cfg.assetsDir = "/mnt/nvme/adoption-os/assets"
	}

	seed := flag.Bool("seed", false, "Seed adoption dates from CSV")
	seedSlugs := flag.Bool("seed-slugs", false, "Seed slugs for existing pets")
	seedVolunteers := flag.Bool("seed-volunteers", false, "Seed active volunteers from mock data")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String(slog.TimeKey, t.Format(time.RFC3339))
				}
			}
			return a
		},
	}))
	slog.SetDefault(logger)

	if err := os.MkdirAll(cfg.assetsDir, 0755); err != nil {
		logger.Error("failed to create assets directory", "path", cfg.assetsDir, "error", err)
		os.Exit(1)
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Could not connect to database.", "error", err)
	} else {
		defer db.Close()
		logger.Info("database connection pool established")

		if err := RunMigrations(db); err != nil {
			logger.Error("migration failure", "error", err)
			os.Exit(1)
		}
		logger.Info("migrations completed")
	}

	if *seed || *seedSlugs || *seedVolunteers {
		if db == nil {
			logger.Error("Cannot seed without database connection")
			os.Exit(1)
		}
		models := data.NewModels(db)

		if *seed {
			logger.Info("Starting Adoption Date Seeding...")
			if err := models.Pets.SeedAdoptionDates(); err != nil {
				logger.Error("Seeding failed", "error", err)
				os.Exit(1)
			}
			logger.Info("Adoption Date Seeding completed successfully")
		}

		if *seedSlugs {
			logger.Info("Starting Pet Slug Seeding...")
			if err := models.Pets.SeedSlugs(); err != nil {
				logger.Error("Pet Slug Seeding failed", "error", err)
				os.Exit(1)
			}
			logger.Info("Pet Slug Seeding completed successfully")
		}

		if *seedVolunteers {
			app := &application{
				config: cfg,
				logger: logger,
				db:     db,
			}
			logger.Info("Starting Volunteer Seeding...")
			app.seedVolunteers()
			logger.Info("Volunteer Seeding completed successfully")
		}
		os.Exit(0)
	}

	app := &application{
		config:   cfg,
		logger:   logger,
		models:   data.NewModels(db),
		mailer:   mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
		notifier: notifier.New(data.NewModels(db), logger),
		db:       db,
	}

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			err := app.models.Applications.ArchiveOldPending(ctx)
			if err != nil {
				app.logger.Error("Background Worker: Failed to archive", "error", err)
			}

			err = app.models.Applications.DeleteDeniedApplications(ctx)
			if err != nil {
				app.logger.Error("Background Worker: Failed to delete denied", "error", err)
			}
			cancel()
		}
	}()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "env", cfg.env, "addr", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Signal Received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exiting")
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
