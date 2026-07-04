package mysql

import "time"

type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
}

type GetOneUserParam struct {
	ID    int
	Name  string
	Email string
}
