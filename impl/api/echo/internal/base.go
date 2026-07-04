package internal

import (
	"context"

	"simple-blog.com/comment/comment"
	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

type Post interface {
	CreatePost(ctx context.Context, post entity.Post) (entity.Post, error)
	ListPost(ctx context.Context, param post.ListPostParam) (post.GetPaginatedPostResponse, error)
	GetPost(ctx context.Context, ID int) (entity.Post, error)
	UpdatePost(ctx context.Context, req post.UpdatePostParam) error
	DeletePost(ctx context.Context, ID int, requesterID int) error
}

type Comment interface {
	AddCommment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	ListComment(ctx context.Context, param comment.ListCommentParam) (comment.GetPaginatedCommentResponse, error)
}

type User interface {
	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
	Register(ctx context.Context, req user.RegisterRequest) error
}

type Sanitizer interface {
	Sanitize(html string) string
}

type Echo struct {
	pst       Post
	cmt       Comment
	usr       User
	sanitizer Sanitizer
}

func NewEcho(pst Post, cmt Comment, usr User, sanitizer Sanitizer) *Echo {
	return &Echo{pst, cmt, usr, sanitizer}
}
