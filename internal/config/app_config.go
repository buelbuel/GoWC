package config

import (
	"os"
	"strings"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pelletier/go-toml"
	"golang.org/x/crypto/acme/autocert"
)

// AppConfig represents the configuration for the application.
type AppConfig struct {
	ServerAddress    string            `toml:"ServerAddress"`
	StaticPaths      map[string]string `toml:"StaticPaths"`
	UseLogger        bool              `toml:"UseLogger"`
	LogOutput        string            `toml:"LogOutput"`
	LogFile          string            `toml:"LogFile"`
	ColorizeLogger   bool              `toml:"ColorizeLogger"`
	UseTLS           bool              `toml:"UseTLS"`
	UseAutoTLS       bool              `toml:"UseAutoTLS"`
	CertFile         string            `toml:"CertFile"`
	KeyFile          string            `toml:"KeyFile"`
	Domain           string            `toml:"Domain"`
	CacheDir         string            `toml:"CacheDir"`
	EnableCORS       bool              `toml:"EnableCORS"`
	CORSAllowOrigins []string          `toml:"CORSAllowOrigins"`
	CORSAllowMethods []string          `toml:"CORSAllowMethods"`
	RateLimit        float64           `toml:"RateLimit"`
	RateBurst        int               `toml:"RateBurst"`
}

// NewAppConfig creates a new instance of AppConfig.
func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}

// SetupStaticFiles sets up the static files for the application.
func (config *AppConfig) SetupStaticFiles(echo *echo.Echo) {
	for route, path := range config.StaticPaths {
		echo.Static(route, path)
	}
}

// SetupMiddleware sets up the middleware for the application.
func (config *AppConfig) SetupMiddleware(echo *echo.Echo) {
	if config.UseLogger {
		loggerConfig := middleware.LoggerConfig{
			Format: "{\n\t\"time\":\"${time_rfc3339}\",\n\t\"id\":\"${id}\",\n\t\"remote_ip\":\"${remote_ip}\",\n\t\"host\":\"${host}\",\n\t\"method\":\"${method}\",\n\t\"uri\":\"${uri}\",\n\t\"status\":${status},\n\t\"error\":\"${error}\",\n\t\"latency\":${latency},\n\t\"latency_human\":\"${latency_human}\",\n\t\"bytes_in\":${bytes_in},\n\t\"bytes_out\":${bytes_out},\n\t\"user_agent\":\"${user_agent}\"\n},\n",
		}
		echo.Logger.SetLevel(log.DEBUG)

		if config.LogOutput == "file" {
			file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				echo.Logger.Fatal(err)
			}

			loggerConfig.Output = file
		} else if config.LogOutput == "stdout" {
			loggerConfig.Output = os.Stdout
		}

		if config.ColorizeLogger {
			loggerConfig.Format = colorizeLogFormat(loggerConfig.Format)
		}

		echo.Use(middleware.LoggerWithConfig(loggerConfig))
	}

	if config.EnableCORS {
		echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.CORSAllowOrigins,
			AllowMethods: config.CORSAllowMethods,
		}))
	}
}

func colorizeLogFormat(format string) string {
	colorMap := map[string]string{
		"time":          "\033[36m", // Cyan
		"id":            "\033[32m", // Green
		"remote_ip":     "\033[32m", // Green
		"host":          "\033[32m", // Green
		"method":        "\033[35m", // Magenta
		"uri":           "\033[35m", // Magenta
		"status":        "\033[35m", // Magenta
		"error":         "\033[31m", // Red
		"latency":       "\033[34m", // Blue
		"latency_human": "\033[34m", // Blue
		"bytes_in":      "\033[34m", // Blue
		"bytes_out":     "\033[34m", // Blue
		"user_agent":    "\033[33m", // Yellow
	}

	for key, color := range colorMap {
		format = strings.Replace(format, "\""+key+"\":", color+"\""+key+"\":\033[0m", -1)
	}

	return format
}

// SetupRenderer sets up the renderer for the application.
func (config *AppConfig) SetupRenderer(echo *echo.Echo) {
	echo.Renderer = utils.NewTemplates()
}

// StartServer starts the server.
func (config *AppConfig) StartServer(echo *echo.Echo) {
	address := config.ServerAddress

	if address == "" {
		address = "localhost:4000"
		echo.Logger.Warn("No server address provided, defaulting to: %s", address)
	}

	if config.UseAutoTLS {
		echo.AutoTLSManager.Cache = autocert.DirCache(config.CacheDir)
		echo.Logger.Fatal(echo.StartAutoTLS(":443"))
	} else if config.UseTLS {
		echo.Logger.Fatal(echo.StartTLS(address, config.CertFile, config.KeyFile))
	} else {
		echo.Logger.Fatal(echo.Start(address))
	}
}
