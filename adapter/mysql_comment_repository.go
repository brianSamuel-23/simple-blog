package adapter

import (
	"context"

	"simple-blog.com/comment/comment"
	"simple-blog.com/entity"
	"simple-blog.com/impl/repo/mysql/mysql"
)

type CommentRepository interface {
	GetPaginatedComment(ctx context.Context, param mysql.GetPaginatedCommentParam) (mysql.GetPaginatedCommentResponse, error)
	SaveComment(ctx context.Context, comment mysql.Comment) (mysql.Comment, error)
}

type mysqlCommentRepository struct {
	repo CommentRepository
}

func NewMysqlCommentRepository(repo CommentRepository) *mysqlCommentRepository {
	return &mysqlCommentRepository{repo}
}

func (r *mysqlCommentRepository) GetPaginated(ctx context.Context, param comment.GetPaginatedCommentParam) (comment.GetPaginatedCommentResponse, error) {
	resp, err := r.repo.GetPaginatedComment(ctx, mysql.GetPaginatedCommentParam{
		PostID:  param.PostID,
		Page:    param.Page,
		PerPage: param.PerPage,
		Order:   param.Order,
		OrderBy: param.OrderBy,
	})
	if err != nil {
		return comment.GetPaginatedCommentResponse{}, err
	}

	comments := make([]entity.Comment, 0, len(resp.Comments))
	for _, c := range resp.Comments {
		comments = append(comments, toEntityComment(c))
	}

	return comment.GetPaginatedCommentResponse{Comments: comments, TotalData: resp.TotalData}, nil
}

func (r *mysqlCommentRepository) Save(ctx context.Context, c entity.Comment) (entity.Comment, error) {
	saved, err := r.repo.SaveComment(ctx, toMysqlComment(c))
	if err != nil {
		return entity.Comment{}, err
	}

	return toEntityComment(saved), nil
}

func toEntityComment(m mysql.Comment) entity.Comment {
	var authorID int
	if m.AuthorID != nil {
		authorID = *m.AuthorID
	}

	return entity.Comment{
		ID:         m.ID,
		PostID:     m.PostID,
		AuthorID:   authorID,
		AuthorName: m.AuthorName,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
	}
}

func toMysqlComment(e entity.Comment) mysql.Comment {
	var authorID *int
	if e.AuthorID != 0 {
		authorID = &e.AuthorID
	}

	return mysql.Comment{
		ID:         e.ID,
		PostID:     e.PostID,
		AuthorID:   authorID,
		AuthorName: e.AuthorName,
		Content:    e.Content,
		CreatedAt:  e.CreatedAt,
	}
}
