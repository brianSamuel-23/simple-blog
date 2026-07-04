package jwt

import (
	"time"

	"simple-blog.com/impl/jwt/internal"
	"simple-blog.com/impl/jwt/jwt"
)

type TokenIssuer interface {
	Issue(claims jwt.TokenClaims) (token string, expiresAt time.Time, err error)
}

func New(secret string, ttl time.Duration) TokenIssuer {
	return internal.NewIssuer(secret, ttl)
}
