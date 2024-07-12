package config

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml"
)

// TLSConfig represents the configuration for TLS.
// It includes the path to the certificate file, the path to the key file, and a flag to enable TLS.
type TLSConfig struct {
	CertFile string `toml:"CertFile"`
	KeyFile  string `toml:"KeyFile"`
	UseTLS   bool   `toml:"UseTLS"`
}

// NewTLSConfig creates a new instance of [TLSConfig] by reading from the config.toml file.
// It returns a pointer to a [TLSConfig] instance and an error if one occurs.
// It reads the config.toml file from the root directory and unmarshals the TOML data into a [TLSConfig] instance.
// If the file is not found or cannot be read, it returns an error.
func NewTLSConfig() (*TLSConfig, error) {
	config := &TLSConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}

// StartServer starts the echo server with TLS if enabled.
// It takes an [echo.Echo] instance and an address string as arguments.
// It returns an error if the server fails to start.
func (config *TLSConfig) StartServer(echo *echo.Echo, address string) error {
	if config.UseTLS {
		return echo.StartTLS(address, config.CertFile, config.KeyFile)
	}
	return echo.Start(address)
}

// StartCustomServer starts the echo server with TLS if enabled using a custom http.Server.
// It takes an [echo.Echo] instance and an address string as arguments.
// It returns an error if the server fails to start.
func (config *TLSConfig) StartCustomServer(echo *echo.Echo, address string) error {
	server := http.Server{
		Addr:      address,
		Handler:   echo,
		TLSConfig: &tls.Config{},
	}
	if config.UseTLS {
		return server.ListenAndServeTLS(config.CertFile, config.KeyFile)
	}
	return server.ListenAndServe()
}
