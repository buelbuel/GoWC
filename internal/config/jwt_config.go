package config

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pelletier/go-toml"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	SecretKey string
}

func NewJwtConfig() *JwtConfig {
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}

	var config struct {
		Server struct {
			JwtSecretKey string `toml:"jwt_secret_key"`
		}
	}

	if err := toml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	return &JwtConfig{
		SecretKey: config.Server.JwtSecretKey,
	}
}

func (claims *JwtCustomClaims) Valid() error {
	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token has expired")
	}
	return nil
}
