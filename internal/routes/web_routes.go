package routes

import (
	config "github.com/buelbuel/gowc/internal/config"
	handlers "github.com/buelbuel/gowc/internal/handlers"
	layers "github.com/buelbuel/gowc/internal/layers"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

// WebRoutes sets up all the web routes for the application.
// It defines both public and protected routes, associating them with their respective handlers.
//
// Parameters:
//   - echo: The Echo instance to register routes on.
//   - state: The application state, which can be used in handlers if needed.
//   - jwtConfig: JWT configuration for authentication.
//   - authHandlers: Handlers for authentication-related operations.
//
// This function separates routes into two main categories:
// 1. Public routes: Accessible without authentication.
// 2. Protected routes: Require authentication to access.
func WebRoutes(
	echo *echo.Echo,
	state *utils.State,
	jwtConfig *config.JwtConfig,
	authHandlers *handlers.AuthHandlers,
) {
	// Public routes
	echo.GET("/", handlers.StartPageHandler)
	echo.GET("/auth", handlers.AuthPageHandler)
	echo.POST("/api/register", authHandlers.RegisterHandler)
	echo.POST("/api/login", authHandlers.LoginHandler)
	echo.GET("/api/logout", authHandlers.LogoutHandler)

	// Protected routes
	app := echo.Group("/app", layers.RequireAuth(jwtConfig))
	app.GET("/dashboard", handlers.DashboardPageHandler)
	app.GET("/profile", handlers.ProfilePageHandler)
}
