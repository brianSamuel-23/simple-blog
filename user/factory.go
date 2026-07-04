package user

import (
	"context"

	"simple-blog.com/adapter"
	"simple-blog.com/impl/hasher"
	"simple-blog.com/impl/jwt"
	"simple-blog.com/impl/repo/mysql"
	"simple-blog.com/user/internal"
	"simple-blog.com/user/user"
)

type User interface {
	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
	Register(ctx context.Context, req user.RegisterRequest) error
}

func New(repo mysql.Repository, hash hasher.Hasher, jwt jwt.TokenIssuer) User {
	usrRepo := adapter.NewMysqlUserRepository(repo)
	tkn := adapter.NewJWTTokenIssuer(jwt)
	return internal.NewUser(usrRepo, hash, tkn)
}
