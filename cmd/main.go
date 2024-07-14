package main

import (
	config "github.com/buelbuel/gowc/internal/config"
	routes "github.com/buelbuel/gowc/internal/routes"
	"github.com/labstack/echo/v4"
)

// main is the entry point for the web application.
func main() {
	echo := echo.New()
	appConfig, err := config.NewAppConfig()
	if err != nil {
		echo.Logger.Fatal(err)
	}
	appConfig.SetupMiddleware(echo)
	appConfig.SetupStaticFiles(echo)
	appConfig.SetupRenderer(echo)

	routes.WebRoutes(echo)

	appConfig.StartServer(echo)
}
