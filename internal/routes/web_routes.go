package routes

import (
	handlers "github.com/buelbuel/gowired/internal/handlers"
	middleware "github.com/buelbuel/gowired/internal/middleware"
	"github.com/labstack/echo/v4"
)

func WebRoutes(echo *echo.Echo) {
	// Public routes
	echo.GET("/", handlers.StartPageHandler)
	echo.GET("/register", handlers.RegisterPageHandler)
	echo.GET("/login", handlers.LoginPageHandler)
	echo.GET("/logout", handlers.LogoutHandler)

	// Authenticated routes
	app := echo.Group("/app", middleware.RequireAuth)
	app.GET("/dashboard", handlers.DashboardPageHandler)
	app.GET("/profile", handlers.ProfilePageHandler)
}
