package internal

import (
	"context"

	"simple-blog.com/comment/comment"
	"simple-blog.com/entity"
	"simple-blog.com/post/post"
	"simple-blog.com/user/user"
)

type CommentRepository interface {
	GetPaginated(ctx context.Context, param comment.GetPaginatedCommentParam) (comment.GetPaginatedCommentResponse, error)
	Save(ctx context.Context, comment entity.Comment) (entity.Comment, error)
}

type UserRepository interface {
	GetOne(ctx context.Context, param user.GetOneUserParam) (entity.User, error)
}

type PostRepository interface {
	GetOne(ctx context.Context, param post.GetOnePostParam) (entity.Post, error)
}

type Comment struct {
	cmtRepo CommentRepository
	usrRepo UserRepository
	pstRepo PostRepository
}

func NewComment(cmtRepo CommentRepository, usrRepo UserRepository, pstRepo PostRepository) *Comment {
	return &Comment{cmtRepo, usrRepo, pstRepo}
}
