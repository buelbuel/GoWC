package handlers

import (
	"net/http"

	models "github.com/buelbuel/gowc/internal/models"
	"github.com/labstack/echo/v4"
)

// UserHandlers represents the handlers for user-related operations.
type UserHandlers struct {
	UserModel models.UserInterface
}

// NewUserHandlers creates a new instance of UserHandlers.
func NewUserHandlers(userModel models.UserInterface) *UserHandlers {
	return &UserHandlers{UserModel: userModel}
}

// CreateUser represents the handler for creating a user.
func (handler *UserHandlers) CreateUser(context echo.Context) error {
	user := new(models.User)
	if err := context.Bind(user); err != nil {
		context.Logger().Errorf("Failed to bind user input: %v", err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := handler.UserModel.CreateUser(user); err != nil {
		context.Logger().Errorf("Failed to create user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	context.Logger().Info("User created successfully with ID: " + user.ID)
	return context.JSON(http.StatusCreated, user)
}

// GetUser represents the handler for getting a user.
func (handler *UserHandlers) GetUser(context echo.Context) error {
	id := context.Param("id")
	user, err := handler.UserModel.GetUser(id)
	if err != nil {
		context.Logger().Errorf("Failed to get user: %v", err)
		return context.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	context.Logger().Info("User found with ID: " + user.ID)
	return context.JSON(http.StatusOK, user)
}

// UpdateUser represents the handler for updating a user.
func (handler *UserHandlers) UpdateUser(context echo.Context) error {
	id := context.Param("id")
	user := new(models.User)
	if err := context.Bind(user); err != nil {
		context.Logger().Errorf("Failed to bind user input: %v", err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	user.ID = id

	if err := handler.UserModel.UpdateUser(user); err != nil {
		context.Logger().Errorf("Failed to update user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	context.Logger().Info("User updated successfully with ID: " + user.ID)
	return context.JSON(http.StatusOK, user)
}

// DeleteUser represents the handler for deleting a user.
func (handler *UserHandlers) DeleteUser(context echo.Context) error {
	id := context.Param("id")
	if err := handler.UserModel.DeleteUser(id); err != nil {
		context.Logger().Errorf("Failed to delete user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	context.Logger().Info("User deleted successfully with ID: " + id)
	return context.NoContent(http.StatusNoContent)
}
