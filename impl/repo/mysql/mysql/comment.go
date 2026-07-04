package mysql

import (
	"time"
)

type Comment struct {
	ID         int
	PostID     int
	AuthorID   *int
	AuthorName string
	Content    string

	CreatedAt time.Time
}

type ListCommentParam struct {
	PostID  int
	Page    int
	PerPage int
	Order   string
	OrderBy string
}

type GetPaginatedCommentParam struct {
	PostID  int
	Page    int
	PerPage int
	Order   string
	OrderBy string
}

type GetPaginatedCommentResponse struct {
	Comments  []Comment
	TotalData int
}
