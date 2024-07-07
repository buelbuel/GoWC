package handlers

import (
	"github.com/labstack/echo/v4"
)

func SaveUser(context echo.Context) error {
	return context.String(200, "User saved")
}

func GetUser(context echo.Context) error {
	return context.String(200, "User retrieved")
}

func UpdateUser(context echo.Context) error {
	return context.String(200, "User updated")
}

func DeleteUser(context echo.Context) error {
	return context.String(200, "User deleted")
}
