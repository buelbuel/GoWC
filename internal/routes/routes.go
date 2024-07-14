package routes

import (
	handlers "github.com/buelbuel/gowc/internal/handlers"
	"github.com/labstack/echo/v4"
)

// WebRoutes sets up all the web routes for the application.
func WebRoutes(echo *echo.Echo) {
	echo.GET("/", handlers.StartPageHandler)
}
