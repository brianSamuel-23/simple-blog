package internal

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"simple-blog.com/impl/repo/mysql/mysql"
	"simple-blog.com/post/post"
)

var postOrderableColumns = map[string]bool{
	"id":          true,
	"title":       true,
	"author_id":   true,
	"author_name": true,
	"created_at":  true,
	"updated_at":  true,
}

func (m *MySQL) SavePost(ctx context.Context, post mysql.Post) (mysql.Post, error) {

	if err := m.db.WithContext(ctx).Save(&post).Error; err != nil {
		return mysql.Post{}, err
	}

	return post, nil
}

func (m *MySQL) DeletePost(ctx context.Context, ID int) error {
	return m.db.WithContext(ctx).Delete(&mysql.Post{}, ID).Error
}

func (m *MySQL) GetPaginatedPost(ctx context.Context, param mysql.GetPaginatedPostParam) (mysql.GetPaginatedPostResponse, error) {
	var rows []mysql.Post
	var total int64

	query := m.db.WithContext(ctx).Model(&mysql.Post{})

	if err := query.Count(&total).Error; err != nil {
		return mysql.GetPaginatedPostResponse{}, err
	}

	if param.OrderBy != "" {
		if order := buildOrderClause(param.OrderBy, param.Order, postOrderableColumns); order != "" {
			query = query.Order(order)
		}
	}

	page, perPage := param.Page, param.PerPage
	if page < 1 {
		page = 1
	}
	if perPage > 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}

	if err := query.Find(&rows).Error; err != nil {
		return mysql.GetPaginatedPostResponse{}, err
	}

	return mysql.GetPaginatedPostResponse{Posts: rows, TotalData: int(total)}, nil
}

func (m *MySQL) GetOnePost(ctx context.Context, param mysql.GetOnePostParam) (mysql.Post, error) {
	var row mysql.Post

	err := m.db.WithContext(ctx).First(&row, param.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return mysql.Post{}, post.ErrPostNotFound
	}
	if err != nil {
		return mysql.Post{}, err
	}

	return row, nil
}
