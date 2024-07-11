// internal/routes/api_routes.go
package routes

import (
	config "github.com/buelbuel/gowc/internal/config"
	handlers "github.com/buelbuel/gowc/internal/handlers"
	layers "github.com/buelbuel/gowc/internal/layers"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(
	echo *echo.Echo,
	state *utils.State,
	jwtConfig *config.JwtConfig,
	userHandlers *handlers.UserHandlers,
	authHandlers *handlers.AuthHandlers,
) {
	// Public routes
	echo.POST("/api/login", authHandlers.LoginHandler)
	echo.POST("/api/register", authHandlers.RegisterHandler)

	// Protected routes
	auth := echo.Group("/api")
	auth.Use(layers.RequireAuth(jwtConfig))
	auth.Use(layers.CheckBlacklist(state))

	auth.GET("/logout", authHandlers.LogoutHandler)
	auth.POST("/users", userHandlers.CreateUser)
	auth.GET("/users/:id", userHandlers.GetUser)
	auth.PUT("/users/:id", userHandlers.UpdateUser)
	auth.DELETE("/users/:id", userHandlers.DeleteUser)
}
