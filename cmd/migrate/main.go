package main

import (
	"database/sql"
	"flag"
	"log"

	config "github.com/buelbuel/gowc/internal/config"
	migrations "github.com/buelbuel/gowc/internal/migrations"
	_ "github.com/lib/pq"
)

// MigrationDirection represents the direction of a migration.
type MigrationDirection string

// Constants for migration directions.
const (
	MigrationDirectionUp   MigrationDirection = "up"
	MigrationDirectionDown MigrationDirection = "down"
)

// main is the entry point for the migration tool.
func main() {
	direction := flag.String("direction", "up", "Migration direction: up or down")
	flag.Parse()

	dbConfig, err := config.NewDBConfig()
	checkError("Failed to load database configuration", err)

	db, err := sql.Open("postgres", dbConfig.DatabaseURL)
	checkError("Failed to connect to database", err)
	defer db.Close()

	runner := migrations.NewMigrationRunner(db)

	migrationsToRun := []migrations.Migration{
		&migrations.CreateUsersTable{},
		// More migrations can be added here
	}

	switch MigrationDirection(*direction) {
	case MigrationDirectionUp:
		err = runner.RunMigrations(migrationsToRun)
		checkError("Failed to run migrations", err)
		log.Println("Migrations applied successfully")
	case MigrationDirectionDown:
		err = runner.RevertMigrations(migrationsToRun)
		checkError("Failed to revert migrations", err)
		log.Println("Migrations reverted successfully")
	default:
		log.Fatalf("Invalid direction: %s. Use 'up' or 'down'.", *direction)
	}
}

// checkError checks if an error occurred and logs it.
func checkError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
