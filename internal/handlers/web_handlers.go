package handlers

import (
	"net/http"

	config "github.com/buelbuel/gowc/internal/config"
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

func AuthPageHandler(context echo.Context) error {
	return renderPage(context, "Auth", "AuthPageContent", "FrontLayout")
}

func LoginFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<register-form></register-form>`)
}

func RegisterFormHandler(context echo.Context) error {
	return context.HTML(http.StatusOK, `<login-form></login-form>`)
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

func LogoutHandler(context echo.Context) error {
	return context.Redirect(http.StatusSeeOther, "/")
}
