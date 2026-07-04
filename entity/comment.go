package entity

import "time"

type Comment struct {
	ID          [16]byte
	PostID      [16]byte
	AuthorID    [16]byte
	AuthorName  string
	ContentHTML string
	ContentRaw  string
	CreatedAt   time.Time
}
