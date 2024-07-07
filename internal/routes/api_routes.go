package routes

import (
	"database/sql"

	handlers "github.com/buelbuel/gowc/internal/handlers"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(echo *echo.Echo, state *utils.State, db *sql.DB) {
	authHandlers := handlers.NewAuthHandlers(state, db)

	echo.POST("/users", handlers.SaveUser)
	echo.GET("/users/:id", handlers.GetUser)
	echo.PUT("/users/:id", handlers.UpdateUser)
	echo.DELETE("/users/:id", handlers.DeleteUser)
	echo.GET("/login", authHandlers.LoginHandler)
	echo.GET("/logout", authHandlers.LogoutHandler)
}
