package main

import (
	"log"

	config "github.com/buelbuel/gowc/internal/config"
	routes "github.com/buelbuel/gowc/internal/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	echo := echo.New()
	appConfig := config.NewAppConfig()
	appConfig.SetupMiddleware(echo)
	appConfig.SetupStaticFiles(echo)
	appConfig.SetupRenderer(echo)
	appState := config.InitializeState()
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConfig.DB.Close()
	jwtConfig := config.NewJwtConfig()
	routes.WebRoutes(echo, appState, jwtConfig, dbConfig.DB)
	routes.ApiRoutes(echo, appState, dbConfig.DB)
	appConfig.StartServer(echo)
}
