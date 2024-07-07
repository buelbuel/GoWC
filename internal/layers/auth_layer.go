package layers

import (
	"github.com/buelbuel/gowc/internal/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RequireAuth(jwtConfig *config.JwtConfig) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtConfig.SecretKey),
	})
}
