package config

import (
	"os"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pelletier/go-toml"
	"golang.org/x/crypto/acme/autocert"
)

type AppConfig struct {
	ServerAddress    string
	StaticPaths      map[string]string
	UseLogger        bool
	UseTLS           bool
	UseAutoTLS       bool
	CertFile         string
	KeyFile          string
	Domain           string
	CacheDir         string
	EnableCORS       bool
	CORSAllowOrigins []string
	CORSAllowMethods []string
}

func NewAppConfig() *AppConfig {
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	var config struct {
		Server struct {
			Address          string
			UseLogger        bool
			UseTLS           bool
			CertFile         string
			KeyFile          string
			UseAutoTLS       bool
			Domain           string
			CacheDir         string
			EnableCORS       bool
			CORSAllowOrigins []string
			CORSAllowMethods []string
		}
		StaticPaths map[string]string
	}

	if err := toml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	return &AppConfig{
		ServerAddress:    config.Server.Address,
		StaticPaths:      config.StaticPaths,
		UseLogger:        config.Server.UseLogger,
		UseTLS:           config.Server.UseTLS,
		CertFile:         config.Server.CertFile,
		KeyFile:          config.Server.KeyFile,
		UseAutoTLS:       config.Server.UseAutoTLS,
		Domain:           config.Server.Domain,
		CacheDir:         config.Server.CacheDir,
		EnableCORS:       config.Server.EnableCORS,
		CORSAllowOrigins: config.Server.CORSAllowOrigins,
		CORSAllowMethods: config.Server.CORSAllowMethods,
	}
}

func (config *AppConfig) SetupStaticFiles(echo *echo.Echo) {
	for route, path := range config.StaticPaths {
		echo.Static(route, path)
	}
}

func (config *AppConfig) SetupMiddleware(echo *echo.Echo) {
	if config.UseLogger {
		echo.Use(middleware.Logger())
	}
	if config.EnableCORS {
		echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.CORSAllowOrigins,
			AllowMethods: config.CORSAllowMethods,
		}))
	}
}

func (config *AppConfig) SetupRenderer(echo *echo.Echo) {
	echo.Renderer = utils.NewTemplates()
}

func (config *AppConfig) StartServer(echo *echo.Echo) {
	if config.UseAutoTLS {
		echo.AutoTLSManager.Cache = autocert.DirCache(config.CacheDir)
		echo.Logger.Fatal(echo.StartAutoTLS(":443"))
	} else if config.UseTLS {
		echo.Logger.Fatal(echo.StartTLS(config.ServerAddress, config.CertFile, config.KeyFile))
	} else {
		echo.Logger.Fatal(echo.Start(config.ServerAddress))
	}
}
