package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuthenticated := false

		if !isAuthenticated {
			return c.Redirect(http.StatusSeeOther, "/login")
		}

		return next(c)
	}
}
