package config

import (
	"io"
	"os"
	"strings"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pelletier/go-toml"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/time/rate"
)

// AppConfig represents the application configuration.
// It contains various settings for server, logging, TLS, CORS, and rate limiting.
// These are mapped with toml tags from which the config is read in [NewAppConfig].
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

// NewAppConfig creates a new instance of [AppConfig] by reading from the config.toml file.
// It reads the config.toml file from the root directory and unmarshals the TOML data into an [AppConfig] instance.
// If the file is not found or cannot be read, it returns an error.
//
// Returns:
//   - *AppConfig: A pointer to the loaded configuration
//   - error: An error if the configuration file cannot be read or parsed
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

// SetupMiddleware configures and sets up the middleware for the [echo.Echo] instance.
// This is where the middleware is set up, such as logging, CORS, and rate limiting.
// It takes a pointer to an [AppConfig] instance and an [echo.Echo] instance as arguments.
func (config *AppConfig) SetupMiddleware(echo *echo.Echo) {
	// Setup logging middleware if enabled
	if config.UseLogger {
		echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: config.getLogFormat(),
			Output: config.getLogOutput(),
		}))
	}

	// Setup CORS middleware if enabled
	if config.EnableCORS {
		echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.CORSAllowOrigins,
			AllowMethods: config.CORSAllowMethods,
		}))
	}

	// Setup rate limiting middleware
	echo.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(config.RateLimit))))
}

// SetupStaticFiles configures static file serving for the [echo.Echo] instance.
// It takes a pointer to an [AppConfig] instance and an [echo.Echo] instance as arguments.
// It sets up static file serving for each route specified in the [AppConfig]'s [StaticPaths] map.
func (config *AppConfig) SetupStaticFiles(echo *echo.Echo) {
	for route, path := range config.StaticPaths {
		echo.Static(route, path)
	}
}

// getLogFormat returns the log format string, optionally colorized.
// It returns the log format string based on the [AppConfig]'s [UseLogger] and [ColorizeLogger] settings.
func (config *AppConfig) getLogFormat() string {
	format := `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n"

	if config.ColorizeLogger {
		return colorizeLogFormat(format)
	}
	return format
}

// getLogOutput returns the appropriate [io.Writer] for logging based on configuration.
// It returns the appropriate [io.Writer] based on the [AppConfig]'s [LogOutput] setting.
// If the [LogOutput] is set to "file", it returns a file writer.
// If the [LogOutput] is set to "stdout", it returns [os.Stdout].
// Otherwise, it returns [os.Stderr].
func (config *AppConfig) getLogOutput() io.Writer {
	switch config.LogOutput {
	case "file":
		file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Warn("Failed to log to file, using default stderr")
			return os.Stderr
		}
		return file
	case "stdout":
		return os.Stdout
	default:
		return os.Stderr
	}
}

// colorizeLogFormat adds ANSI color codes to the log format string.
// It takes a string as an argument and returns the colorized string.
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
// It takes an [echo.Echo] instance as an argument and sets its Renderer field.
func (config *AppConfig) SetupRenderer(echo *echo.Echo) {
	echo.Renderer = utils.NewTemplates()
}

// StartServer starts the server based on the configuration.
// It handles different scenarios like AutoTLS, TLS, and standard HTTP.
// It takes an [echo.Echo] instance as an argument.
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
