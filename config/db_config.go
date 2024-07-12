package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // Blank import for PostgreSQL driver
	"github.com/pelletier/go-toml"
)

// DBConfig creates a new instance of [DatabaseConfig] with the provided parameters.
//
// Parameters:
//   - Url: The hostname of the database server
//   - MaxConns: The port number of the database server
//   - MaxIdleConns: The username for database authentication
//   - ConnMaxLifetime: The password for database authentication
//
// Returns:
//   - *DBConfig: A pointer to the created [DBConfig] instance
type DBConfig struct {
	Database struct {
		URL             string        `toml:"URL"`
		MaxConns        int           `toml:"MaxConns"`
		MaxIdleConns    int           `toml:"MaxIdleConns"`
		ConnMaxLifetime time.Duration `toml:"ConnMaxLifetime"`
	} `toml:"Database"`
	DB *sql.DB `toml:"-"`
}

// NewDBConfig creates a new DBConfig instance by reading the configuration from a toml file.
// It returns a pointer to a [DBConfig] instance and an error if one occurs.
// It reads the config.toml file from the root directory and unmarshals the TOML data into a [DBConfig] instance.
// If the file is not found or cannot be read, it returns an error.
func NewDBConfig() (*DBConfig, error) {
	config := &DBConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}
	log.Printf("Config file contents: %s", string(configFile))

	if err := toml.Unmarshal(configFile, config); err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		return nil, err
	}

	log.Printf("Loaded database URL: %s", config.Database.URL)
	log.Printf("Full config: %+v", config)

	return config, nil
}

// Initialize initializes the database connection and sets connection pool settings.
// It opens a database connection using the provided [DBConfig]'s [DatabaseURL] and sets the maximum number of open connections,
// maximum idle connections, and connection maximum lifetime.
func (config *DBConfig) Initialize() error {
	db, err := sql.Open("postgres", config.Database.URL)
	if err != nil {
		return err
	}

	config.DB = db
	config.DB.SetMaxOpenConns(config.Database.MaxConns)
	config.DB.SetMaxIdleConns(config.Database.MaxIdleConns)
	config.DB.SetConnMaxLifetime(config.Database.ConnMaxLifetime)

	return nil
}
