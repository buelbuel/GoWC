package handlers

import (
	"net/http"

	config "github.com/buelbuel/gowired/internal/config"
	"github.com/labstack/echo/v4"
)

func renderPage(context echo.Context, title, contentBlock, layout string) error {
	data := map[string]interface{}{
		"Title":        title,
		"ContentBlock": contentBlock,
		"echoContext":  context,
		"Menu":         config.GetMenu(layout),
	}
	return context.Render(http.StatusOK, layout, data)
}

func StartPageHandler(context echo.Context) error {
	return renderPage(context, "Start", "startPageContent", "FrontLayout")
}

func DashboardPageHandler(context echo.Context) error {
	return renderPage(context, "Dashboard", "dashboardPageContent", "AppLayout")
}

func ProfilePageHandler(context echo.Context) error {
	return renderPage(context, "Profile", "profilePageContent", "AppLayout")
}

func RegisterPageHandler(context echo.Context) error {
	return renderPage(context, "Register", "registerPageContent", "FrontLayout")
}

func LoginPageHandler(context echo.Context) error {
	return renderPage(context, "Login", "loginPageContent", "FrontLayout")
}

func LogoutHandler(context echo.Context) error {
	return context.Redirect(http.StatusSeeOther, "/")
}
