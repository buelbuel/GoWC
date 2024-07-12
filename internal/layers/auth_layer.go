package layers

import (
	"time"

	config "github.com/buelbuel/gowc/internal/config"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// RequireAuth requires a valid JWT token to access the route.
// It returns an echo.MiddlewareFunc that can be used as a middleware for echo.
func RequireAuth(jwtConfig *config.JwtConfig) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtConfig.SecretKey),
	})
}

// CheckBlacklist checks if the token is blacklisted.
// It returns an echo.MiddlewareFunc that can be used as a middleware for echo.
func CheckBlacklist(state *utils.State) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			token := context.Get("user").(*jwt.Token)
			blacklistedTokens := state.GetState()["blacklistedTokens"].(map[string]time.Time)

			if _, blacklisted := blacklistedTokens[token.Raw]; blacklisted {
				context.Logger().Error("Token blacklisted")
				return echo.ErrUnauthorized
			}

			context.Logger().Info("Token not blacklisted")
			return next(context)
		}
	}
}
