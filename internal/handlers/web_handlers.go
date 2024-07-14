package handlers

import (
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
)

// StartPageHandler displays the start page.
func StartPageHandler(context echo.Context) error {
	return utils.RenderPage(context, "Start", "Description for Start Page", "startPageContent", "FrontLayout")
}
