package routes

import (
	config "github.com/buelbuel/gowc/config"
	handlers "github.com/buelbuel/gowc/handlers"
	layers "github.com/buelbuel/gowc/layers"
	utils "github.com/buelbuel/gowc/utils"
	"github.com/labstack/echo/v4"
)

// APIRoutes sets up all the API routes for the application.
// It defines both public and protected API endpoints, associating them with their respective handlers.
//
// Parameters:
//   - echo: The Echo instance to register routes on.
//   - state: The application state, which can be used in handlers if needed.
//   - jwtConfig: JWT configuration for authentication.
//   - userHandlers: Handlers for user-related operations.
//   - authHandlers: Handlers for authentication-related operations.
//
// This function separates routes into two main categories:
// 1. Public routes: Accessible without authentication.
// 2. Protected routes: Require authentication to access.
func APIRoutes(
	echo *echo.Echo,
	state *utils.State,
	jwtConfig *config.JwtConfig,
	userHandlers *handlers.UserHandlers,
	authHandlers *handlers.AuthHandlers,
) {
	// Public routes
	// These API endpoints are accessible without authentication
	echo.POST("/api/login", authHandlers.LoginHandler)
	echo.POST("/api/register", authHandlers.RegisterHandler)

	// Protected routes
	// These API endpoints require authentication to access
	auth := echo.Group("/api")
	auth.Use(layers.RequireAuth(jwtConfig))
	auth.Use(layers.CheckBlacklist(state))

	auth.GET("/logout", authHandlers.LogoutHandler)
	auth.POST("/users", userHandlers.CreateUser)
	auth.GET("/users/:id", userHandlers.GetUser)
	auth.PUT("/users/:id", userHandlers.UpdateUser)
	auth.DELETE("/users/:id", userHandlers.DeleteUser)
}
