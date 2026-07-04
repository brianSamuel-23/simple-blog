package comment

import "simple-blog.com/entity"

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
	Comments  []entity.Comment
	TotalData int
}
