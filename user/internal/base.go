package internal

import (
	"context"
	"time"

	"simple-blog.com/entity"
	"simple-blog.com/user/user"
)

type UserRepository interface {
	GetOne(ctx context.Context, param user.GetOneUserParam) (entity.User, error)
	Create(ctx context.Context, usr entity.User) (entity.User, error)
}

type PasswordHasher interface {
	Hash(plaintext string) (string, error)
	Compare(hash, plaintext string) error
}

type TokenIssuer interface {
	Issue(claims user.TokenClaims) (token string, expiresAt time.Time, err error)
}

type User struct {
	usrRepo UserRepository
	hasher  PasswordHasher
	tokens  TokenIssuer
}

func NewUser(usrRepo UserRepository, hasher PasswordHasher, tokens TokenIssuer) *User {
	return &User{usrRepo, hasher, tokens}
}
