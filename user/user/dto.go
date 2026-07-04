package user

import "time"

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginResponse struct {
	LoginID     int
	AccessToken string
	ExpiresAt   time.Time
}

type GetOneUserParam struct {
	ID    int
	Name  string
	Email string
}

type TokenClaims struct {
	UserID int
	Name   string
}
