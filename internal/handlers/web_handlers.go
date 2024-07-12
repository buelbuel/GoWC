package handlers

import (
	"net/http"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

// AuthPageHandler displays the authentication page.
func AuthPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Auth", "AuthPageContent", "FrontLayout")
}

// LoginFormHandler handles the login form.
func LoginFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<register-form></register-form>`)
}

// RegisterFormHandler handles the registration form.
func RegisterFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<login-form></login-form>`)
}

// StartPageHandler displays the start page.
func StartPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Start", "startPageContent", "FrontLayout")
}

// DashboardPageHandler displays the dashboard page.
func DashboardPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Dashboard", "dashboardPageContent", "AppLayout")
}

// ProfilePageHandler displays the profile page.
func ProfilePageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Profile", "profilePageContent", "AppLayout")
}
