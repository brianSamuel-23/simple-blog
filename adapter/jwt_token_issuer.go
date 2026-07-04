package adapter

import (
	"time"

	"simple-blog.com/impl/jwt/jwt"
	"simple-blog.com/user/user"
)

type TokenIssuer interface {
	Issue(claims jwt.TokenClaims) (token string, expiresAt time.Time, err error)
}

type JWTTokenIssuer struct {
	tkn TokenIssuer
}

func NewJWTTokenIssuer(tkn TokenIssuer) *JWTTokenIssuer {
	return &JWTTokenIssuer{tkn}
}

func (j *JWTTokenIssuer) Issue(claims user.TokenClaims) (token string, expiresAt time.Time, err error) {
	return j.tkn.Issue(jwt.TokenClaims{
		UserID: claims.UserID,
		Name:   claims.Name,
	})
}
