package internal

import (
	"context"

	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

type PostRepository interface {
	Save(ctx context.Context, post entity.Post) (entity.Post, error)
	Delete(ctx context.Context, ID int) error
	GetPaginated(ctx context.Context, param post.GetPaginatedPostParam) (post.GetPaginatedPostResponse, error)
	GetOne(ctx context.Context, param post.GetOnePostParam) (entity.Post, error)
}

type UserRepository interface {
	GetOne(ctx context.Context, param user.GetOneUserParam) (entity.User, error)
}

type Post struct {
	postRepo PostRepository
	usrRepo  UserRepository
}

func NewPost(postRepo PostRepository, usrRepo UserRepository) *Post {
	return &Post{postRepo, usrRepo}
}
