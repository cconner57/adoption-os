package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
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
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
	mailer mailer.Mailer // Add Mailer dependency
	db     *sql.DB       // Add DB connection for HealthCheck
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found (relying on system env vars)")
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

	flag.Parse()

	// Initialize the logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create the DB connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Printf("WARNING: Could not connect to database: %v. Running in CSV-fallback mode.", err)
		// db will be nil or unusable, handled in models
	} else {
		defer db.Close()
		logger.Printf("database connection pool established")
	}

	// 👇 INITIALIZE THE APP INSTANCE
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
		db:     db,
	}

	// Start the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(), // Call the routes method
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
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
	logger.Println("Shutdown Signal Received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}

	// db.Close() will be called by defer at top of main()
	logger.Println("Server exiting")
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
