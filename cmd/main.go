package main

import (
	routes "github.com/buelbuel/gowc/internal/routes"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echo := echo.New()
	echo.Static("/css", "public/css")
	echo.Static("/js", "public/js")
	echo.Static("/images", "public/images")
	echo.Use(middleware.Logger())
	echo.Renderer = utils.NewTemplates()
	routes.WebRoutes(echo)
	echo.Logger.Fatal(echo.Start("localhost:4000"))
}
