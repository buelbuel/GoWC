package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq" // Blank import for PostgreSQL driver
	"github.com/pelletier/go-toml"
)

// DBConfig represents the configuration for the database.
type DBConfig struct {
	DatabaseURL     string        `toml:"Url"`
	MaxConns        int           `toml:"MaxConns"`
	MaxIdleConns    int           `toml:"MaxIdleConns"`
	ConnMaxLifetime time.Duration `toml:"ConnMaxLifetime"`
	DB              *sql.DB       `toml:"-"`
}

// NewDBConfig creates a new DBConfig instance by reading the configuration from a toml file.
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
