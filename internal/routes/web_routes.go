package routes

import (
	"database/sql"

	config "github.com/buelbuel/gowc/internal/config"
	handlers "github.com/buelbuel/gowc/internal/handlers"
	layers "github.com/buelbuel/gowc/internal/layers"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

func WebRoutes(
	echo *echo.Echo,
	state *utils.State,
	jwtConfig *config.JwtConfig,
	db *sql.DB,
) {
	authHandlers := handlers.NewAuthHandlers(state, db)

	// Public routes
	echo.GET("/", handlers.StartPageHandler)
	echo.GET("/auth", handlers.AuthPageHandler)
	echo.GET("/login-form", handlers.LoginFormHandler)
	echo.GET("/register-form", handlers.RegisterFormHandler)
	echo.POST("/register", authHandlers.RegisterHandler)
	echo.POST("/login", authHandlers.LoginHandler)
	echo.GET("/logout", authHandlers.LogoutHandler)

	// Authenticated routes
	app := echo.Group("/app", layers.RequireAuth(jwtConfig))
	app.GET("/dashboard", handlers.DashboardPageHandler)
	app.GET("/profile", handlers.ProfilePageHandler)
}
