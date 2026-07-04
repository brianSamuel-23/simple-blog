package internal

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwttoken "simple-blog.com/impl/jwt/jwt"
)

type Issuer struct {
	secret []byte
	ttl    time.Duration
}

func NewIssuer(secret string, ttl time.Duration) *Issuer {
	return &Issuer{secret: []byte(secret), ttl: ttl}
}

func (i *Issuer) Issue(claims jwttoken.TokenClaims) (string, time.Time, error) {
	expiresAt := time.Now().Add(i.ttl)

	registeredClaims := jwt.MapClaims{
		"sub":  strconv.Itoa(claims.UserID),
		"iss":  "simple-blog-v1",
		"name": claims.Name,
		"iat":  jwt.NewNumericDate(time.Now()),
		"exp":  jwt.NewNumericDate(expiresAt),
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims).SignedString(i.secret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expiresAt, nil
}
