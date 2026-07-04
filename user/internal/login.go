package internal

import (
	"context"
	"errors"

	"simple-blog.com/user/user"
)

func (u *User) Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	usr, err := u.usrRepo.GetOne(ctx, user.GetOneUserParam{Email: req.Email})
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return user.LoginResponse{}, user.ErrInvalidCredentials
		}
		return user.LoginResponse{}, err
	}

	if err := u.hasher.Compare(usr.PasswordHash, req.Password); err != nil {
		return user.LoginResponse{}, user.ErrInvalidCredentials
	}

	token, expiresAt, err := u.tokens.Issue(user.TokenClaims{UserID: usr.ID})
	if err != nil {
		return user.LoginResponse{}, err
	}

	return user.LoginResponse{
		LoginID:     usr.ID,
		AccessToken: token,
		ExpiresAt:   expiresAt,
	}, nil
}
