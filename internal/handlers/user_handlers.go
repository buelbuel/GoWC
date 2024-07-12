package handlers

import (
	"net/http"

	models "github.com/buelbuel/gowc/internal/models"
	"github.com/labstack/echo/v4"
)

// UserHandlers handles HTTP requests related to user operations.
// It provides methods for creating, retrieving, updating, and deleting users.
//
// The struct holds references to:
//   - UserModel: Interface for user-related database operations
type UserHandlers struct {
	UserModel models.UserInterface
}

// NewUserHandlers creates a new instance of [UserHandlers].
// It takes a [models.UserInterface] as a parameter to handle user-related database operations.
// It returns a pointer to a [UserHandlers] instance.
func NewUserHandlers(userModel models.UserInterface) *UserHandlers {
	return &UserHandlers{UserModel: userModel}
}

// CreateUser handles the creation of a new user.
// It binds the incoming JSON request to a [models.User] struct,
// creates the user in the database, and returns the created user or an error.
//
// Returns:
//   - 201 Created with the user data on success
//   - 400 Bad Request if the input is invalid
//   - 500 Internal Server Error if there's a failure in user creation
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

// GetUser handles the HTTP GET request to retrieve a user by ID.
// It extracts the user ID from the request parameters,
// fetches the user from the database, and returns the user data.
//
// Returns:
//   - 200 OK with the user data on success
//   - 404 Not Found if the user doesn't exist
//   - 500 Internal Server Error if there's a failure in fetching the user
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

// UpdateUser handles updating an existing user.
// It binds the incoming JSON request to a [models.User] struct,
// updates the user in the database, and returns the updated user data.
//
// Returns:
//   - 200 OK with the updated user data on success
//   - 400 Bad Request if the input is invalid
//   - 500 Internal Server Error if there's a failure in updating the user
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
// It extracts the user ID from the request parameters and
// deletes the corresponding user from the database.
//
// Returns:
//   - 204 No Content on successful deletion
//   - 500 Internal Server Error if there's a failure in deleting the user
func (handler *UserHandlers) DeleteUser(context echo.Context) error {
	id := context.Param("id")
	if err := handler.UserModel.DeleteUser(id); err != nil {
		context.Logger().Errorf("Failed to delete user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	context.Logger().Info("User deleted successfully with ID: " + id)
	return context.NoContent(http.StatusNoContent)
}
