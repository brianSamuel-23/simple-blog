package entity

import "time"

type Post struct {
	ID         int
	Title      string
	Content    string
	AuthorID   int
	AuthorName string

	Timestamps
	DeletedAt time.Time
}
