package adapter

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"simple-blog.com/entity"
	"simple-blog.com/impl/repo/mysql/mysql"
	"simple-blog.com/post/post"
)

type PostRepository interface {
	SavePost(ctx context.Context, post mysql.Post) (mysql.Post, error)
	DeletePost(ctx context.Context, ID int) error
	GetPaginatedPost(ctx context.Context, param mysql.GetPaginatedPostParam) (mysql.GetPaginatedPostResponse, error)
	GetOnePost(ctx context.Context, param mysql.GetOnePostParam) (mysql.Post, error)
}

type mysqlPostRepository struct {
	repo PostRepository
}

func NewMysqlPostRepository(repo PostRepository) *mysqlPostRepository {
	return &mysqlPostRepository{repo}
}

func (r *mysqlPostRepository) Save(ctx context.Context, p entity.Post) (entity.Post, error) {
	saved, err := r.repo.SavePost(ctx, toMysqlPost(p))
	if err != nil {
		return entity.Post{}, err
	}

	return toEntityPost(saved), nil
}

func (r *mysqlPostRepository) Delete(ctx context.Context, ID int) error {
	return r.repo.DeletePost(ctx, ID)
}

func (r *mysqlPostRepository) GetPaginated(ctx context.Context, param post.GetPaginatedPostParam) (post.GetPaginatedPostResponse, error) {
	resp, err := r.repo.GetPaginatedPost(ctx, mysql.GetPaginatedPostParam{
		Page:    param.Page,
		PerPage: param.PerPage,
		Order:   param.Order,
		OrderBy: param.OrderBy,
	})
	if err != nil {
		return post.GetPaginatedPostResponse{}, err
	}

	posts := make([]entity.Post, 0, len(resp.Posts))
	for _, p := range resp.Posts {
		posts = append(posts, toEntityPost(p))
	}

	return post.GetPaginatedPostResponse{Posts: posts, TotalData: resp.TotalData}, nil
}

func (r *mysqlPostRepository) GetOne(ctx context.Context, param post.GetOnePostParam) (entity.Post, error) {
	row, err := r.repo.GetOnePost(ctx, mysql.GetOnePostParam{ID: param.ID})
	if err != nil && errors.Is(gorm.ErrRecordNotFound, err) {
		return entity.Post{}, post.ErrPostNotFound
	} else if err != nil {
		return entity.Post{}, err
	}

	return toEntityPost(row), nil
}

func toEntityPost(m mysql.Post) entity.Post {
	return entity.Post{
		ID:         m.ID,
		Title:      m.Title,
		Content:    m.Content,
		AuthorID:   m.AuthorID,
		AuthorName: m.AuthorName,
		Timestamps: entity.Timestamps{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		DeletedAt: m.DeletedAt.Time,
	}
}

func toMysqlPost(e entity.Post) mysql.Post {
	return mysql.Post{
		ID:         e.ID,
		Title:      e.Title,
		Content:    e.Content,
		AuthorID:   e.AuthorID,
		AuthorName: e.AuthorName,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		DeletedAt:  gorm.DeletedAt{Time: e.DeletedAt},
	}
}
