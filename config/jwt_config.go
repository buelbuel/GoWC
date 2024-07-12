package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pelletier/go-toml"
)

// JwtCustomClaims represents the custom claims for the JWT.
// It embeds the standard claims and adds a Name field.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// JwtConfig represents the configuration for JSON Web Tokens (JWT).
// It contains settings for token signing, expiration, and claims.
type JwtConfig struct {
	SecretKey      string        `toml:"JwtSecretKey"`
	ExpirationTime time.Duration `toml:"JwtExpirationTime"`
}

// NewJwtConfig creates a new instance of [JwtConfig] with the provided signing key and expiration time.
//
// Parameters:
//   - signingKey: The key used for signing the JWT
//   - expirationTime: The duration for which the token is valid
//
// Returns:
//   - *JwtConfig: A pointer to the new JwtConfig instance
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

// Valid validates the custom claims in a JWT token.
// It checks if the token has expired.
func (claims *JwtCustomClaims) Valid() error {
	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return jwt.ErrTokenExpired
	}
	return nil
}
