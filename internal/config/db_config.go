package config

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

type DBConfig struct {
	DatabaseURL     string
	DB              *sql.DB
	MaxConns        int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewDBConfig() (*DBConfig, error) {
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	var config struct {
		Database struct {
			Url             string
			MaxConns        int
			MaxIdleConns    int
			ConnMaxLifetime string
		}
	}

	if err := toml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", config.Database.Url)
	if err != nil {
		return nil, err
	}

	connMaxLifetime, err := time.ParseDuration(config.Database.ConnMaxLifetime)
	if err != nil {
		return nil, err
	}

	dbConfig := &DBConfig{
		DatabaseURL:     config.Database.Url,
		DB:              db,
		MaxConns:        config.Database.MaxConns,
		MaxIdleConns:    config.Database.MaxIdleConns,
		ConnMaxLifetime: connMaxLifetime,
	}

	db.SetMaxOpenConns(dbConfig.MaxConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	return dbConfig, nil
}
