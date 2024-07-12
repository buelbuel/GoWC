package migrate

import (
	"flag"
	"log"

	config "github.com/buelbuel/gowc/config"
	migrations "github.com/buelbuel/gowc/migrations"
)

// MigrationDirection represents the direction of a migration.
type MigrationDirection string

// Constants for migration directions.
const (
	MigrationDirectionUp   MigrationDirection = "up"
	MigrationDirectionDown MigrationDirection = "down"
)

// Run is the entry point for the migration tool.
func Run() {
	direction := flag.String("direction", "up", "Migration direction: up or down")
	flag.Parse()

	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	if dbConfig.Database.URL == "" {
		log.Fatalf("Database URL is empty. Please check your config.toml file.")
	}

	err = dbConfig.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer dbConfig.DB.Close()

	runner := migrations.NewMigrationRunner(dbConfig.DB)

	migrationsToRun := []migrations.Migration{
		&migrations.CreateUsersTable{},
		// Add more migrations here as needed
	}

	switch MigrationDirection(*direction) {
	case MigrationDirectionUp:
		err = runner.RunMigrations(migrationsToRun)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations applied successfully")
	case MigrationDirectionDown:
		err = runner.RevertMigrations(migrationsToRun)
		if err != nil {
			log.Fatalf("Failed to revert migrations: %v", err)
		}
		log.Println("Migrations reverted successfully")
	default:
		log.Fatalf("Invalid direction: %s. Use 'up' or 'down'.", *direction)
	}
}
