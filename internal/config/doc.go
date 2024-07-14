// Package config provides configuration structures and utilities for the application.
//
// This package includes configurations for:
//   - Application settings ([AppConfig])
//
// The configurations are typically loaded from a TOML file using the
// [github.com/pelletier/go-toml] package. The [NewAppConfig] function
// is responsible for reading and parsing the configuration file.
//
// Usage:
//
//	config, err := config.NewAppConfig()
//	if err != nil {
//	    log.Fatalf("Failed to load configuration: %v", err)
//	}
//
// The loaded configuration can then be used to set up various parts of the application,
// such as the server, database connection, and JWT authentication.
package config
