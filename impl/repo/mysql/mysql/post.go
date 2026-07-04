package mysql

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	Content    string
	AuthorID   int
	AuthorName string
	CreatedAt  time.Time `gorm:"<-:create"`
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type GetPaginatedPostParam struct {
	Page    int
	PerPage int
	Order   string
	OrderBy string
}

type GetOnePostParam struct {
	ID int
}

type GetPaginatedPostResponse struct {
	Posts     []Post
	TotalData int
}
