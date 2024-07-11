package layers

import (
	"time"

	config "github.com/buelbuel/gowc/internal/config"
	utils "github.com/buelbuel/gowc/internal/utils"
	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RequireAuth(jwtConfig *config.JwtConfig) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtConfig.SecretKey),
	})
}

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
