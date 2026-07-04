package user

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found!")
	ErrEmailAlreadyExists = errors.New("email is already registered")
	ErrWeakPassword       = errors.New("password is too weak: combine at least 3 of uppercase, lowercase, digit, and special characters, and avoid common or personal passwords")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
