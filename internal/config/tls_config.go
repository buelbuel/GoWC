package config

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml"
)

// TLSConfig represents the configuration for TLS.
type TLSConfig struct {
	CertFile string `toml:"CertFile"`
	KeyFile  string `toml:"KeyFile"`
	UseTLS   bool   `toml:"UseTLS"`
}

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

func (config *TLSConfig) StartServer(echo *echo.Echo, address string) error {
	if config.UseTLS {
		return echo.StartTLS(address, config.CertFile, config.KeyFile)
	}
	return echo.Start(address)
}

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
