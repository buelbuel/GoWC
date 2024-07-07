package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/buelbuel/gowc/internal/config"
	"github.com/buelbuel/gowc/internal/models"
	"github.com/buelbuel/gowc/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	State *utils.State
	DB    *sql.DB
}

func NewAuthHandlers(state *utils.State, db *sql.DB) *AuthHandlers {
	return &AuthHandlers{State: state, DB: db}
}

func (handlers *AuthHandlers) RegisterHandler(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := user.Create(handlers.DB); err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
	})
}

func (handlers *AuthHandlers) LoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.GetUserByUsername(handlers.DB, username)
	if err != nil {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return echo.ErrUnauthorized
	}

	claims := &config.JwtCustomClaims{
		Name:  user.Username,
		Admin: user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenEncoded, err := token.SignedString([]byte("your_secret_key_here"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenEncoded,
	})
}

func (handlers *AuthHandlers) LogoutHandler(context echo.Context) error {
	// TODO: Implement logout logic here
	return context.String(http.StatusOK, "Logged out")
}

func (handlers *AuthHandlers) Accessible(context echo.Context) error {
	return context.String(http.StatusOK, "Accessible")
}

func (handlers *AuthHandlers) Restricted(context echo.Context) error {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	name := claims.Name
	return context.String(http.StatusOK, "Welcome "+name+"!")
}
