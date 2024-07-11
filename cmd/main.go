package main

import (
	config "github.com/buelbuel/gowc/internal/config"
	handlers "github.com/buelbuel/gowc/internal/handlers"
	models "github.com/buelbuel/gowc/internal/models"
	routes "github.com/buelbuel/gowc/internal/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	echo := echo.New()

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

	appConfig.SetupMiddleware(echo)
	appConfig.SetupStaticFiles(echo)
	appConfig.SetupRenderer(echo)
	appState := stateConfig.InitializeState()

	userModel := models.NewUserModel(dbConfig.DB, echo.Logger)
	userHandlers := handlers.NewUserHandlers(userModel)
	authHandlers := handlers.NewAuthHandlers(appState, userModel, jwtConfig)

	routes.WebRoutes(echo, appState, jwtConfig, userHandlers, authHandlers)
	routes.ApiRoutes(echo, appState, jwtConfig, userHandlers, authHandlers)

	appConfig.StartServer(echo)
}
