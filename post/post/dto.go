package post

import "simple-blog.com/entity"

type ListPostParam struct {
	Page    int
	PerPage int
	Order   string
	OrderBy string
}

type ListPostResponse struct {
	posts     []entity.Post
	TotalData int
}

type UpdatePostParam struct {
	ID          int
	Title       string
	Content     string
	RequesterID int
}

type GetPaginatedPostParam struct {
	Page    int
	PerPage int
	Order   string
	OrderBy string
}

type GetPaginatedPostResponse struct {
	Posts     []entity.Post
	TotalData int
}

type GetOnePostParam struct {
	ID int
}
