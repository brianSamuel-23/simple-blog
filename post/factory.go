package post

import (
	"context"

	"simple-blog.com/adapter"
	"simple-blog.com/entity"
	"simple-blog.com/impl/repo/mysql"
	"simple-blog.com/post/internal"
	"simple-blog.com/post/post"
)

type Post interface {
	CreatePost(ctx context.Context, post entity.Post) (entity.Post, error)
	ListPost(ctx context.Context, param post.ListPostParam) (post.GetPaginatedPostResponse, error)
	GetPost(ctx context.Context, ID int) (entity.Post, error)
	UpdatePost(ctx context.Context, req post.UpdatePostParam) error
	DeletePost(ctx context.Context, ID int, requesterID int) error
}

func New(repo mysql.Repository) Post {
	postRepo := adapter.NewMysqlPostRepository(repo)
	usrRepo := adapter.NewMysqlUserRepository(repo)
	return internal.NewPost(postRepo, usrRepo)
}
