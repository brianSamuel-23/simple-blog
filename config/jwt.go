package config

import (
	"time"

	"simple-blog.com/impl/jwt"
)

func NewJWT() jwt.TokenIssuer {
	ttlMinutes := getEnvInt("JWT_TTL_MINUTES", 60)

	return jwt.New(jwtSecret(), time.Duration(ttlMinutes)*time.Minute)
}

// jwtSecret returns the key used to both sign and validate tokens, so the
// echo-jwt middleware in NewEcho() stays in sync with the issuer above.
func jwtSecret() string {
	return getEnv("JWT_SECRET", "change-me-in-development")
}
