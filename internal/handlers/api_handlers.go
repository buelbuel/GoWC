package handlers

import (
	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	return c.String(200, "User saved")
}

func GetUser(c echo.Context) error {
	return c.String(200, "User retrieved")
}

func UpdateUser(c echo.Context) error {
	return c.String(200, "User updated")
}

func DeleteUser(c echo.Context) error {
	return c.String(200, "User deleted")
}
