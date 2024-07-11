package config

import (
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
	SecretKey      string        `toml:"JwtSecretKey"`
	ExpirationTime time.Duration `toml:"JwtExpirationTime"`
}

func NewJwtConfig() (*JwtConfig, error) {
	config := &JwtConfig{}
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (claims *JwtCustomClaims) Valid() error {
	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return jwt.ErrTokenExpired
	}
	return nil
}
