package comment

import (
	"context"

	"simple-blog.com/adapter"
	"simple-blog.com/comment/comment"
	"simple-blog.com/comment/internal"
	"simple-blog.com/entity"
	"simple-blog.com/impl/repo/mysql"
)

type Comment interface {
	AddCommment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	ListComment(ctx context.Context, param comment.ListCommentParam) (comment.GetPaginatedCommentResponse, error)
}

func New(repo mysql.Repository) Comment {
	cmtRepo := adapter.NewMysqlCommentRepository(repo)
	usrRepo := adapter.NewMysqlUserRepository(repo)
	pstRepo := adapter.NewMysqlPostRepository(repo)
	return internal.NewComment(cmtRepo, usrRepo, pstRepo)
}
