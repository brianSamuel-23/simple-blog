package internal

import (
	"context"
	"errors"

	"simple-blog.com/entity"
	"simple-blog.com/user/user"
)

func (u *User) Register(ctx context.Context, req user.RegisterRequest) error {
	if !isPasswordStrong(req.Password, req.Name, req.Email) {
		return user.ErrWeakPassword
	}

	if _, err := u.usrRepo.GetOne(ctx, user.GetOneUserParam{Email: req.Email}); err == nil {
		return user.ErrEmailAlreadyExists
	} else if !errors.Is(err, user.ErrUserNotFound) {
		return err
	}

	passwordHash, err := u.hasher.Hash(req.Password)
	if err != nil {
		return err
	}

	if _, err := u.usrRepo.Create(ctx, entity.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}); err != nil {
		return err
	}

	return nil
}
