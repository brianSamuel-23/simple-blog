package mysql

import (
	"context"

	"gorm.io/gorm"

	"simple-blog.com/impl/repo/mysql/internal"
	"simple-blog.com/impl/repo/mysql/mysql"
)

type PostRepository interface {
	SavePost(ctx context.Context, post mysql.Post) (mysql.Post, error)
	DeletePost(ctx context.Context, ID int) error
	GetPaginatedPost(ctx context.Context, param mysql.GetPaginatedPostParam) (mysql.GetPaginatedPostResponse, error)
	GetOnePost(ctx context.Context, param mysql.GetOnePostParam) (mysql.Post, error)
}

type UserRepository interface {
	GetOneUser(ctx context.Context, param mysql.GetOneUserParam) (mysql.User, error)
	CreateUser(ctx context.Context, row mysql.User) (mysql.User, error)
}

type CommentRepository interface {
	GetPaginatedComment(ctx context.Context, param mysql.GetPaginatedCommentParam) (mysql.GetPaginatedCommentResponse, error)
	SaveComment(ctx context.Context, comment mysql.Comment) (mysql.Comment, error)
}

type Repository interface {
	PostRepository
	UserRepository
	CommentRepository
}

func New(db *gorm.DB) Repository {
	return internal.NewMySQL(db)
}
