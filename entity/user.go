package entity

import "time"

type User struct {
	ID           [16]byte
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
