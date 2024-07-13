package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	handlers "github.com/buelbuel/gowc/handlers"
	models "github.com/buelbuel/gowc/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	echo := echo.New()

	mockUserModel := new(MockUserModel)
	mockUserModel.On("GetUser", "1").Return(&models.User{ID: "1", Username: "testuser", Email: "test@example.com"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	context := echo.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	handlers := &handlers.UserHandlers{UserModel: mockUserModel}

	if assert.NoError(t, handlers.GetUser(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"ID":"1","Email":"test@example.com","Username":"testuser","Password":"","Admin":false,"IsActive":false,"IsDeleted":false,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":"0001-01-01T00:00:00Z"}`+"\n", rec.Body.String())
	}

	mockUserModel.AssertCalled(t, "GetUser", "1")
}

func TestUpdateUser(t *testing.T) {
	echo := echo.New()

	mockUserModel := new(MockUserModel)
	mockUserModel.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(`{"id":"1","username":"testuser","email":"test@example.com","password":"testpassword"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	context := echo.NewContext(req, rec)
	handlers := &handlers.UserHandlers{UserModel: mockUserModel}

	if assert.NoError(t, handlers.UpdateUser(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"ID":"","Email":"test@example.com","Username":"testuser","Password":"testpassword","Admin":false,"IsActive":false,"IsDeleted":false,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":"0001-01-01T00:00:00Z"}`+"\n", rec.Body.String())
	}

	mockUserModel.AssertCalled(t, "UpdateUser", mock.AnythingOfType("*models.User"))
}

func TestDeleteUser(t *testing.T) {
	echo := echo.New()

	mockUserModel := new(MockUserModel)
	mockUserModel.On("DeleteUser", "1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec := httptest.NewRecorder()
	context := echo.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")
	handlers := &handlers.UserHandlers{UserModel: mockUserModel}

	if assert.NoError(t, handlers.DeleteUser(context)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}

	mockUserModel.AssertCalled(t, "DeleteUser", "1")
}

func TestGetUserByEmail(t *testing.T) {
	echo := echo.New()

	mockUserModel := new(MockUserModel)
	mockUserModel.On("GetUserByEmail", "test@example.com").Return(&models.User{ID: "1", Username: "testuser", Email: "test@texample.com"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	context := echo.NewContext(req, rec)
	context.SetPath("/users/:email")
	context.SetParamNames("email")
	context.SetParamValues("test@example.com")
	handlers := &handlers.UserHandlers{UserModel: mockUserModel}

	if assert.NoError(t, handlers.GetUserByEmail(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"ID":"1","Email":"test@texample.com","Username":"testuser","Password":"","Admin":false,"IsActive":false,"IsDeleted":false,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":"0001-01-01T00:00:00Z"}`+"\n", rec.Body.String())
	}

	mockUserModel.AssertCalled(t, "GetUserByEmail", "test@example.com")
}
