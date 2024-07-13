package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	config "github.com/buelbuel/gowc/config"
	handlers "github.com/buelbuel/gowc/handlers"
	models "github.com/buelbuel/gowc/models"
	utils "github.com/buelbuel/gowc/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterHandler(t *testing.T) {
	e := echo.New()

	mockUserModel := new(MockUserModel)
	mockUserModel.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)

	jwtConfig := &config.JwtConfig{}
	state := &utils.State{}
	handlers := &handlers.AuthHandlers{State: state, UserModel: mockUserModel, JwtConfig: jwtConfig}

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"username":"testuser","email":"test@example.com","password":"password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	if assert.NoError(t, handlers.RegisterHandler(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "User registered successfully")
	}

	mockUserModel.AssertCalled(t, "CreateUser", mock.AnythingOfType("*models.User"))
}

func TestLoginHandler(t *testing.T) {
	e := echo.New()

	mockUserModel := new(MockUserModel)
	validPasswordHash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockUserModel.On("GetUserByEmail", "test@example.com").Return(&models.User{ID: "1", Username: "testuser", Email: "test@example.com", Password: string(validPasswordHash)}, nil)

	jwtConfig := &config.JwtConfig{}
	state := &utils.State{}
	handlers := &handlers.AuthHandlers{State: state, UserModel: mockUserModel, JwtConfig: jwtConfig}

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com","password":"password123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	if assert.NoError(t, handlers.LoginHandler(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
	}

	mockUserModel.AssertCalled(t, "GetUserByEmail", "test@example.com")
}
