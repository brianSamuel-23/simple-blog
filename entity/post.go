package entity

import "time"

type Post struct {
	ID        [16]byte
	Title     string
	Content   string
	AuthorID  [16]byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
