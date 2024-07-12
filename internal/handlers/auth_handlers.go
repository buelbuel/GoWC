package handlers

import (
	"fmt"
	"net/http"
	"time"

	config "github.com/buelbuel/gowc/internal/config"
	models "github.com/buelbuel/gowc/internal/models"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandlers represents the handlers for authentication operations.
// It contains methods for user registration, login, logout, and token refresh.
//
// The struct holds references to:
//   - State: Application state for managing things like blacklisted tokens
//   - UserModel: Interface for user-related database operations
//   - JwtConfig: Configuration for JWT token generation and validation
type AuthHandlers struct {
	State     *utils.State
	UserModel models.UserInterface
	JwtConfig *config.JwtConfig
}

// NewAuthHandlers creates a new instance of [AuthHandlers].
// It takes a pointer to the application state, a user model interface, and a JWT configuration.
// It returns a pointer to an [AuthHandlers] instance.
func NewAuthHandlers(state *utils.State, userModel models.UserInterface, jwtConfig *config.JwtConfig) *AuthHandlers {
	return &AuthHandlers{
		State:     state,
		UserModel: userModel,
		JwtConfig: jwtConfig,
	}
}

// RegisterHandler processes user registration requests.
//
// It performs the following steps:
//  1. Binds the incoming JSON request to a struct containing username, email, and password.
//  2. Validates that all required fields are present.
//  3. Hashes the provided password using bcrypt.
//  4. Creates a new user record in the database.
//  5. Returns a JSON response indicating success or failure.
//
// Error cases:
//   - Returns 400 Bad Request if the input is invalid or missing required fields.
//   - Returns 500 Internal Server Error if there's a failure in password hashing or user creation.
//
// On success, it returns a 201 Created status with a success message.
func (handler *AuthHandlers) RegisterHandler(context echo.Context) error {
	var registerData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := context.Bind(&registerData); err != nil {
		context.Logger().Errorf("Failed to parse registration data: %v", err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	context.Logger().Infof("Received registration data: %+v", registerData)

	if registerData.Email == "" || registerData.Username == "" || registerData.Password == "" {
		context.Logger().Error("Email, Username, and Password are required.")
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Email, Username, and Password are required"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		context.Logger().Errorf("Failed to hash password: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	context.Logger().Infof("Hashed password length: %d", len(hashedPassword))

	user := &models.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: string(hashedPassword),
	}

	if err := handler.UserModel.CreateUser(user); err != nil {
		context.Logger().Errorf("Failed to create user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	context.Logger().Info("User registered successfully with ID: " + user.ID)
	return context.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

// LoginHandler processes user login requests.
//
// It performs the following steps:
//  1. Binds the incoming JSON request to a struct containing email and password.
//  2. Retrieves the user from the database using the provided email.
//  3. Compares the provided password with the stored hashed password.
//  4. Generates a JWT token upon successful authentication.
//  5. Returns the JWT token in the response.
//
// Error cases:
//   - Returns 400 Bad Request if the input is invalid.
//   - Returns 401 Unauthorized if the credentials are invalid.
//   - Returns 500 Internal Server Error if there's a failure in token generation.
//
// On success, it returns a 200 OK status with the JWT token.
func (handler *AuthHandlers) LoginHandler(context echo.Context) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := context.Bind(&loginData); err != nil {
		context.Logger().Errorf("Failed to parse login data: %v", err)
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	context.Logger().Infof("Received login data: %+v", loginData)

	user, err := handler.UserModel.GetUserByEmail(loginData.Email)
	if err != nil {
		context.Logger().Errorf("Failed to get user: %v", err)
		return context.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	context.Logger().Infof("Retrieved user: %+v", user)
	context.Logger().Infof("Stored hashed password length: %d", len(user.Password))

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		context.Logger().Errorf("Failed to compare password: %v", err)
		return context.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	claims := &config.JwtCustomClaims{
		Name:  user.Username,
		Admin: user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(handler.JwtConfig.ExpirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenEncoded, err := token.SignedString([]byte(handler.JwtConfig.SecretKey))
	if err != nil {
		context.Logger().Errorf("Failed to generate token: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	context.Logger().Info("User logged in successfully")
	return context.JSON(http.StatusOK, map[string]string{"token": tokenEncoded})
}

// RefreshToken handles the token refresh process.
//
// It performs the following steps:
//  1. Parses and validates the provided refresh token.
//  2. Extracts claims from the valid token.
//  3. Creates a new token with updated expiration time.
//  4. Returns the new token in the response.
//
// Error cases:
//   - Returns 401 Unauthorized if the refresh token is invalid or expired.
//   - Returns 500 Internal Server Error if there's a failure in generating the new token.
//
// On success, it returns a 200 OK status with the new JWT token.
func (handler *AuthHandlers) RefreshToken(context echo.Context) error {
	refreshToken := context.FormValue("refresh_token")
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			context.Logger().Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		context.Logger().Info("Refresh token verified")
		return []byte(handler.JwtConfig.SecretKey), nil
	})

	if err != nil {
		context.Logger().Errorf("Failed to parse token: %v", err)
		return context.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		newClaims := &config.JwtCustomClaims{
			Name:  claims["name"].(string),
			Admin: claims["admin"].(bool),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(handler.JwtConfig.ExpirationTime)),
			},
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		newTokenEncoded, err := newToken.SignedString([]byte(handler.JwtConfig.SecretKey))
		if err != nil {
			context.Logger().Errorf("Failed to generate new token: %v", err)
			return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate new token"})
		}

		context.Logger().Info("Token refreshed successfully")
		return context.JSON(http.StatusOK, map[string]string{"token": newTokenEncoded})
	}

	context.Logger().Error("Invalid refresh token")
	return context.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid refresh token"})
}

// LogoutHandler processes user logout requests.
//
// It performs the following steps:
//  1. Retrieves the current JWT token from the request context.
//  2. Adds the token to a blacklist with an expiration time.
//  3. Updates the application state with the blacklisted token.
//
// This method effectively invalidates the token for future use.
//
// On success, it returns a 200 OK status with a logout confirmation message.
func (handler *AuthHandlers) LogoutHandler(context echo.Context) error {
	token := context.Get("user").(*jwt.Token)

	blacklistedTokens := handler.State.GetState()["blacklistedTokens"].(map[string]time.Time)
	blacklistedTokens[token.Raw] = time.Now().Add(handler.JwtConfig.ExpirationTime)

	handler.State.GetState()["blacklistedTokens"] = blacklistedTokens

	context.Logger().Info("Logged out successfully")
	return context.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

// Accessible is a handler for testing publicly accessible endpoints.
//
// This handler doesn't require authentication and always returns a successful response.
// It's typically used for health checks or public API endpoints.
//
// Returns a 200 OK status with a simple text message.
func (handler *AuthHandlers) Accessible(context echo.Context) error {
	context.Logger().Info("Accessible")
	return context.String(http.StatusOK, "Accessible")
}

// Restricted is a handler for testing authenticated endpoints.
//
// This handler requires a valid JWT token to access. It demonstrates how to:
//  1. Retrieve the authenticated user's information from the JWT claims.
//  2. Use the user's information in the response.
//
// Returns a 200 OK status with a personalized welcome message.
func (handler *AuthHandlers) Restricted(context echo.Context) error {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustomClaims)
	name := claims.Name

	context.Logger().Info("Restricted")
	return context.String(http.StatusOK, "Welcome "+name+"!")
}
