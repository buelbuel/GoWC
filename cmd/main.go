package main

import (
	config "github.com/buelbuel/gowc/config"
	handlers "github.com/buelbuel/gowc/handlers"
	models "github.com/buelbuel/gowc/models"
	routes "github.com/buelbuel/gowc/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"flag"
	"os"

	migrate "github.com/buelbuel/gowc/cmd/migrate"
)

// main is the entry point for the web application.
func main() {
	migrateFlag := flag.String("migrate", "", "Run database migrations: up or down")
	flag.Parse()

	if *migrateFlag != "" {
		os.Args = []string{os.Args[0], "-direction", *migrateFlag}
		migrate.Run()
		return
	}

	echo := echo.New()

	// Load the application configuration.
	appConfig, err := config.NewAppConfig()
	if err != nil {
		echo.Logger.Fatal(err)
	}
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		echo.Logger.Fatal(err)
	}
	jwtConfig, err := config.NewJwtConfig()
	if err != nil {
		echo.Logger.Fatal(err)
	}
	stateConfig, err := config.NewStateConfig()
	if err != nil {
		echo.Logger.Fatal(err)
	}

	// Setup the application middleware.
	appConfig.SetupMiddleware(echo)
	appConfig.SetupStaticFiles(echo)
	appConfig.SetupRenderer(echo)
	appState := stateConfig.InitializeState()

	// Setup the application models and handlers.
	userModel := models.NewUserModel(dbConfig.DB, echo.Logger)
	userHandlers := handlers.NewUserHandlers(userModel)
	authHandlers := handlers.NewAuthHandlers(appState, userModel, jwtConfig)

	// Setup the application routes.
	routes.WebRoutes(echo, appState, jwtConfig, authHandlers)
	routes.APIRoutes(echo, appState, jwtConfig, userHandlers, authHandlers)

	appConfig.StartServer(echo)
}
