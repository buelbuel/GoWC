package migrations

import (
	"database/sql"
	"log"
)

// Migration represents a database migration.
// It has methods for applying and reverting the migration.
type Migration interface {
	Up() string
	Down() string
	Name() string
}

// MigrationRunner is a runner for database migrations.
// It has methods for creating the migrations table, checking if a migration has been applied,
// applying and reverting migrations, and running and reverting all migrations.
type MigrationRunner struct {
	db *sql.DB
}

// NewMigrationRunner creates a new MigrationRunner.
// It takes a pointer to a [sql.DB] instance as an argument and returns a pointer to a [MigrationRunner] instance.
// It returns a pointer to a [MigrationRunner] instance and an error if one occurs.
// It returns an error if the database connection is not successful.
func NewMigrationRunner(db *sql.DB) *MigrationRunner {
	return &MigrationRunner{db: db}
}

// CreateMigrationsTable creates the migrations table.
// It executes a SQL query to create the migrations table if it doesn't exist.
// It returns an error if the query fails.
func (runner *MigrationRunner) CreateMigrationsTable() error {
	_, err := runner.db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL UNIQUE,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	return err
}

// HasMigrationBeenApplied checks if a migration has been applied.
// It takes a string as an argument and returns a boolean and an error.
// It queries the migrations table to check if a migration with the given name has been applied.
// It returns true if the migration has been applied and nil if the query is successful.
// It returns an error if the query fails.
func (runner *MigrationRunner) HasMigrationBeenApplied(name string) (bool, error) {
	var count int
	err := runner.db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = $1", name).Scan(&count)
	return count > 0, err
}

// ApplyMigration applies a migration.
// It takes a [Migration] instance as an argument and returns an error.
// It begins a transaction, executes the migration's up migration SQL query,
// and inserts the migration's name into the migrations table.
// It returns an error if the transaction fails to begin, the up migration query fails,
// or the insertion into the migrations table fails.
func (runner *MigrationRunner) ApplyMigration(migration Migration) error {
	tx, err := runner.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(migration.Up())
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO migrations (name) VALUES ($1)", migration.Name())
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RevertMigration reverts a migration.
// It takes a [Migration] instance as an argument and returns an error.
// It begins a transaction, executes the migration's down migration SQL query,
// and deletes the migration's name from the migrations table.
// It returns an error if the transaction fails to begin, the down migration query fails,
// or the deletion from the migrations table fails.
func (runner *MigrationRunner) RevertMigration(migration Migration) error {
	tx, err := runner.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(migration.Down())
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM migrations WHERE name = $1", migration.Name())
	if err != nil {
		return err
	}

	return tx.Commit()
}

// RunMigrations runs the migrations.
// It takes a slice of [Migration] instances as an argument and returns an error.
// It creates the migrations table if it doesn't exist, then iterates over the migrations slice,
// checking if each migration has been applied. If it hasn't, it applies the migration.
// It returns an error if the migrations table creation fails, or if any of the migrations fails to apply.
func (runner *MigrationRunner) RunMigrations(migrations []Migration) error {
	err := runner.CreateMigrationsTable()
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		applied, err := runner.HasMigrationBeenApplied(migration.Name())
		if err != nil {
			return err
		}

		if !applied {
			log.Printf("Applying migration: %s", migration.Name())
			err = runner.ApplyMigration(migration)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Skipping migration: %s (already applied)", migration.Name())
		}
	}

	return nil
}

// RevertMigrations reverts the migrations.
// It takes a slice of [Migration] instances as an argument and returns an error.
// It creates the migrations table if it doesn't exist, then iterates over the migrations slice in reverse order,
// checking if each migration has been applied. If it has, it reverts the migration.
// It returns an error if the migrations table creation fails, or if any of the migrations fails to revert.
func (runner *MigrationRunner) RevertMigrations(migrations []Migration) error {
	err := runner.CreateMigrationsTable()
	if err != nil {
		return err
	}

	for i := len(migrations) - 1; i >= 0; i-- {
		migration := migrations[i]
		applied, err := runner.HasMigrationBeenApplied(migration.Name())
		if err != nil {
			return err
		}

		if applied {
			log.Printf("Reverting migration: %s", migration.Name())
			err = runner.RevertMigration(migration)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Skipping migration: %s (not applied)", migration.Name())
		}
	}

	return nil
}
