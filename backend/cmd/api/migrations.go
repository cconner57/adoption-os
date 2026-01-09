package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	// 1. Create schema_migrations table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		id SERIAL PRIMARY KEY,
		version TEXT NOT NULL UNIQUE,
		run_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// 2. Read all migration files
	migDir := "./migrations"
	entries, err := os.ReadDir(migDir)
	if err != nil {
		// Just warn if directory missing (might be intended in some envs, though unlikely here)
		fmt.Printf("Warning: could not read migrations directory: %v\n", err)
		return nil
	}

	var filenames []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			filenames = append(filenames, e.Name())
		}
	}
	sort.Strings(filenames)

	// 3. Iterate and run if needed
	for _, filename := range filenames {
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version=$1)", filename).Scan(&exists)
		if err != nil {
			return fmt.Errorf("failed to check migration status for %s: %w", filename, err)
		}

		if !exists {
			fmt.Printf("Migrating: %s\n", filename)

			content, err := os.ReadFile(filepath.Join(migDir, filename))
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", filename, err)
			}

			// Run the migration in a transaction
			tx, err := db.Begin()
			if err != nil {
				return err
			}

			if _, err := tx.Exec(string(content)); err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to execute migration %s: %w", filename, err)
			}

			// Record version
			if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", filename); err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to record migration version %s: %w", filename, err)
			}

			if err := tx.Commit(); err != nil {
				return err
			}
		} else {
			// fmt.Printf("Skipping: %s (already run)\n", filename)
		}
	}

	return nil
}
