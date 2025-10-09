package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5"
)

// Migration represents a database migration
type Migration struct {
	Version     string
	Description string
	UpSQL       string
	DownSQL     string
}

// RunMigrations runs all pending migrations
func RunMigrations() error {
	ctx := context.Background()

	// Create migrations table if it doesn't exist
	if err := createMigrationsTable(ctx); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get list of applied migrations
	appliedMigrations, err := getAppliedMigrations(ctx)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Get all available migrations
	availableMigrations := getAvailableMigrations()

	// Run pending migrations
	for _, migration := range availableMigrations {
		if !isMigrationApplied(appliedMigrations, migration.Version) {
			log.Printf("Running migration %s: %s", migration.Version, migration.Description)
			
			if err := runMigration(ctx, migration); err != nil {
				return fmt.Errorf("failed to run migration %s: %w", migration.Version, err)
			}
			
			log.Printf("Migration %s completed successfully", migration.Version)
		}
	}

	return nil
}

// createMigrationsTable creates the migrations tracking table
func createMigrationsTable(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)
	`
	
	_, err := DB.Exec(ctx, query)
	return err
}

// getAppliedMigrations returns a list of applied migration versions
func getAppliedMigrations(ctx context.Context) ([]string, error) {
	query := `SELECT version FROM schema_migrations ORDER BY version`
	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []string
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}

	return versions, rows.Err()
}

// isMigrationApplied checks if a migration has been applied
func isMigrationApplied(appliedMigrations []string, version string) bool {
	for _, applied := range appliedMigrations {
		if applied == version {
			return true
		}
	}
	return false
}

// runMigration executes a single migration
func runMigration(ctx context.Context, migration Migration) error {
	tx, err := DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Execute the migration SQL
	if _, err := tx.Exec(ctx, migration.UpSQL); err != nil {
		return fmt.Errorf("failed to execute migration SQL: %w", err)
	}

	// Record the migration as applied
	recordQuery := `INSERT INTO schema_migrations (version) VALUES ($1)`
	if _, err := tx.Exec(ctx, recordQuery, migration.Version); err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	return tx.Commit(ctx)
}

// getAvailableMigrations returns all available migrations
func getAvailableMigrations() []Migration {
	return []Migration{
		{
			Version:     "001_initial_schema",
			Description: "Create initial database schema",
			UpSQL:       getInitialSchemaSQL(),
			DownSQL:     getInitialSchemaDownSQL(),
		},
	}
}

// getInitialSchemaSQL returns the initial schema SQL
func getInitialSchemaSQL() string {
	// Read the schema.sql file
	schemaPath := filepath.Join("database", "schema.sql")
	schemaSQL, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Printf("Warning: Could not read schema.sql file: %v", err)
		// Return a minimal schema if file can't be read
		return `
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
			CREATE TABLE IF NOT EXISTS users (
				user_id TEXT PRIMARY KEY DEFAULT uuid_generate_v4()::text,
				name TEXT NOT NULL,
				email TEXT UNIQUE NOT NULL,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
		`
	}
	return string(schemaSQL)
}

// getInitialSchemaDownSQL returns the SQL to rollback the initial schema
func getInitialSchemaDownSQL() string {
	return `
		DROP TABLE IF EXISTS game_players CASCADE;
		DROP TABLE IF EXISTS games CASCADE;
		DROP TABLE IF EXISTS user_sports CASCADE;
		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS sports CASCADE;
		DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;
	`
}

// RollbackLastMigration rolls back the last applied migration
func RollbackLastMigration() error {
	ctx := context.Background()

	// Get the last applied migration
	query := `SELECT version FROM schema_migrations ORDER BY applied_at DESC LIMIT 1`
	var lastVersion string
	err := DB.QueryRow(ctx, query).Scan(&lastVersion)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("no migrations to rollback")
		}
		return fmt.Errorf("failed to get last migration: %w", err)
	}

	// Find the migration details
	availableMigrations := getAvailableMigrations()
	var migration *Migration
	for _, m := range availableMigrations {
		if m.Version == lastVersion {
			migration = &m
			break
		}
	}

	if migration == nil {
		return fmt.Errorf("migration %s not found", lastVersion)
	}

	log.Printf("Rolling back migration %s: %s", migration.Version, migration.Description)

	tx, err := DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Execute the rollback SQL
	if _, err := tx.Exec(ctx, migration.DownSQL); err != nil {
		return fmt.Errorf("failed to execute rollback SQL: %w", err)
	}

	// Remove the migration record
	recordQuery := `DELETE FROM schema_migrations WHERE version = $1`
	if _, err := tx.Exec(ctx, recordQuery, migration.Version); err != nil {
		return fmt.Errorf("failed to remove migration record: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit rollback: %w", err)
	}

	log.Printf("Migration %s rolled back successfully", migration.Version)
	return nil
}

// GetMigrationStatus returns the status of all migrations
func GetMigrationStatus() (map[string]time.Time, error) {
	ctx := context.Background()
	
	query := `SELECT version, applied_at FROM schema_migrations ORDER BY applied_at`
	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	status := make(map[string]time.Time)
	for rows.Next() {
		var version string
		var appliedAt time.Time
		if err := rows.Scan(&version, &appliedAt); err != nil {
			return nil, err
		}
		status[version] = appliedAt
	}

	return status, rows.Err()
}
