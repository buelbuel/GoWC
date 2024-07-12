package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq" // Blank import for PostgreSQL driver
	"github.com/pelletier/go-toml"
)

// NewDatabaseConfig creates a new instance of [DatabaseConfig] with the provided parameters.
//
// Parameters:
//   - host: The hostname of the database server
//   - port: The port number of the database server
//   - user: The username for database authentication
//   - password: The password for database authentication
//   - dbName: The name of the database to connect to
//
// Returns:
//   - *DatabaseConfig: A pointer to the new DatabaseConfig instance
type DBConfig struct {
	DatabaseURL     string        `toml:"Url"`
	MaxConns        int           `toml:"MaxConns"`
	MaxIdleConns    int           `toml:"MaxIdleConns"`
	ConnMaxLifetime time.Duration `toml:"ConnMaxLifetime"`
	DB              *sql.DB       `toml:"-"`
}

// NewDBConfig creates a new DBConfig instance by reading the configuration from a toml file.
// It returns a pointer to a [DBConfig] instance and an error if one occurs.
// It reads the config.toml file from the root directory and unmarshals the TOML data into a [DBConfig] instance.
// If the file is not found or cannot be read, it returns an error.
func NewDBConfig() (*DBConfig, error) {
	config := &DBConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}

// Initialize initializes the database connection and sets connection pool settings.
// It opens a database connection using the provided [DBConfig]'s [DatabaseURL] and sets the maximum number of open connections,
// maximum idle connections, and connection maximum lifetime.
func (config *DBConfig) Initialize() error {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return err
	}

	config.DB = db
	config.DB.SetMaxOpenConns(config.MaxConns)
	config.DB.SetMaxIdleConns(config.MaxIdleConns)
	config.DB.SetConnMaxLifetime(config.ConnMaxLifetime)

	return nil
}
