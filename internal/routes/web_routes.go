package routes

import (
	handlers "github.com/buelbuel/gowc/internal/handlers"
	middleware "github.com/buelbuel/gowc/internal/middleware"
	"github.com/labstack/echo/v4"
)

func WebRoutes(echo *echo.Echo) {
	// Public routes
	echo.GET("/", handlers.StartPageHandler)
	echo.GET("/auth", handlers.AuthPageHandler)
	echo.GET("/login-form", handlers.LoginFormHandler)
	echo.GET("/register-form", handlers.RegisterFormHandler)
	echo.GET("/logout", handlers.LogoutHandler)

	// Authenticated routes
	app := echo.Group("/app", middleware.RequireAuth)
	app.GET("/dashboard", handlers.DashboardPageHandler)
	app.GET("/profile", handlers.ProfilePageHandler)
}
