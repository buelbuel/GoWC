package routes

import (
	handlers "github.com/buelbuel/gowc/internal/handlers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(echo *echo.Echo) {
	echo.POST("/users", handlers.SaveUser)
	echo.GET("/users/:id", handlers.GetUser)
	echo.PUT("/users/:id", handlers.UpdateUser)
	echo.DELETE("/users/:id", handlers.DeleteUser)
}
