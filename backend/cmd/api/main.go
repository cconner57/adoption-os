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
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"strconv" // Helper for port parsing

	"github.com/cconner57/adoption-os/backend/internal/data"
	"github.com/cconner57/adoption-os/backend/internal/mailer" // Import mailer
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
	config config
	logger *slog.Logger
	models data.Models
	mailer mailer.Mailer // Add Mailer dependency
	db     *sql.DB       // Add DB connection for HealthCheck
}

func main() {
	err := godotenv.Load()
	if err != nil {
		// keeping standard log here for pre-setup, or could use fmt
		fmt.Println("Note: .env file not found (relying on system env vars)")
	}

	var cfg config

	// Flags for command line configuration
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// DB
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	// SMTP
	flag.StringVar(&cfg.smtp.host, "smtp-host", os.Getenv("SMTP_SERVER"), "SMTP host")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT")) // Simple fallback for env reading
	flag.IntVar(&cfg.smtp.port, "smtp-port", port, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv("SMTP_LOGIN"), "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv("SMTP_KEY"), "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", os.Getenv("SMTP_SENDER"), "SMTP sender email")

	flag.StringVar(&cfg.assetsDir, "assets-dir", os.Getenv("ASSETS_DIR"), "Directory for storing uploaded assets")
	if cfg.assetsDir == "" {
		cfg.assetsDir = "/mnt/nvme/adoption-os/assets" // Default for production (Raspberry Pi NVMe)
	}

	seed := flag.Bool("seed", false, "Seed adoption dates from CSV")
	seedSlugs := flag.Bool("seed-slugs", false, "Seed slugs for existing pets")
	seedVolunteers := flag.Bool("seed-volunteers", false, "Seed active volunteers from mock data")

	flag.Parse()

	// Initialize the structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			// Ensure timestamp is RFC3339
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String(slog.TimeKey, t.Format(time.RFC3339))
				}
			}
			return a
		},
	}))

	// Set as global logger (optional, but good practice)
	slog.SetDefault(logger)

	// Ensure assets directory exists
	if err := os.MkdirAll(cfg.assetsDir, 0755); err != nil {
		logger.Error("failed to create assets directory", "path", cfg.assetsDir, "error", err)
		os.Exit(1)
	}

	// Create the DB connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Could not connect to database. Running in CSV-fallback mode.", "error", err)
		// db will be nil or unusable, handled in models
	} else {
		defer db.Close()
		logger.Info("database connection pool established")

		// Run Migrations
		if err := RunMigrations(db); err != nil {
			// If migrations fail, we should probably stop the server
			logger.Error("migration failure", "error", err)
			os.Exit(1)
		}
		logger.Info("migrations completed")
		logger.Info("migrations completed")
	}

	// 3. Check Seed Flag
	if *seed || *seedSlugs || *seedVolunteers {
		if db == nil {
			logger.Error("Cannot seed without database connection")
			os.Exit(1)
		}
		models := data.NewModels(db) // Initialize models slightly early for seeding

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
			// We need an app instance to call seedVolunteers, or pass DB to it.
			// seedVolunteers is a method on *application.
			// Let's create a temporary app instance or refactor seedVolunteers to take DB.
			// For simplicity, let's create the app instance here since we have everything except mailer,
			// but seedVolunteers likely only needs DB and logger.

			// Actually, seedVolunteers uses app.db.
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

	// 👇 INITIALIZE THE APP INSTANCE
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
		db:     db,
	}

	// Start Background Worker for Retention Policy
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			err := app.models.Applications.ArchiveOldPending(ctx)
			if err != nil {
				app.logger.Error("Background Worker: Failed to archive old pending applications", "error", err)
			}
			cancel()
		}
	}()

	// Start the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(), // Call the routes method
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError), // Adapter for http.Server
	}

	logger.Info("starting server", "env", cfg.env, "addr", srv.Addr)

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen", "error", err)
			os.Exit(1) // Fatal equivalent
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Signal Received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown", "error", err)
		os.Exit(1)
	}

	// db.Close() will be called by defer at top of main()
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
