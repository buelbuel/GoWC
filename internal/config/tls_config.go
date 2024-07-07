package config

import (
	"crypto/tls"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TLSConfig struct {
	CertFile string
	KeyFile  string
	UseTLS   bool
}

func NewTLSConfig() *TLSConfig {
	return &TLSConfig{
		CertFile: "cert.pem",
		KeyFile:  "key.pem",
		UseTLS:   true,
	}
}

func (config *TLSConfig) StartServer(echo *echo.Echo, address string) {
	if config.UseTLS {
		echo.Logger.Fatal(echo.StartTLS(address, config.CertFile, config.KeyFile))
	} else {
		echo.Logger.Fatal(echo.Start(address))
	}
}

func (config *TLSConfig) StartCustomServer(echo *echo.Echo, address string) {
	server := http.Server{
		Addr:      address,
		Handler:   echo,
		TLSConfig: &tls.Config{},
	}
	echo.Logger.Fatal(server.ListenAndServeTLS(config.CertFile, config.KeyFile))
}
