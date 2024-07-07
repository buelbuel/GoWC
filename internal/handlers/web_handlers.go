package handlers

import (
	"net/http"

	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Auth", "AuthPageContent", "FrontLayout")
}

func LoginFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<register-form></register-form>`)
}

func RegisterFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<login-form></login-form>`)
}

func StartPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Start", "startPageContent", "FrontLayout")
}

func DashboardPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Dashboard", "dashboardPageContent", "AppLayout")
}

func ProfilePageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Profile", "profilePageContent", "AppLayout")
}
