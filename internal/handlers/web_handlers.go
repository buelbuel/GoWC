package handlers

import (
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

// AuthPageHandler displays the authentication page.
func AuthPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Login or Sign up", "Description for Auth Page", "AuthPageContent", "FrontLayout")
}

// StartPageHandler displays the start page.
func StartPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Start", "Description for Start Page", "startPageContent", "FrontLayout")
}

// DashboardPageHandler displays the dashboard page.
func DashboardPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Dashboard", "Description for Dashboard Page", "dashboardPageContent", "AppLayout")
}

// ProfilePageHandler displays the profile page.
func ProfilePageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Profile", "Description for Profile Page", "profilePageContent", "AppLayout")
}
