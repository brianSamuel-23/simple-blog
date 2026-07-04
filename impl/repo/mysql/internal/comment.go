package internal

import (
	"context"

	"simple-blog.com/impl/repo/mysql/mysql"
)

var commentOrderableColumns = map[string]bool{
	"id":          true,
	"post_id":     true,
	"author_id":   true,
	"author_name": true,
	"created_at":  true,
}

func (m *MySQL) GetPaginatedComment(ctx context.Context, param mysql.GetPaginatedCommentParam) (mysql.GetPaginatedCommentResponse, error) {
	var rows []mysql.Comment
	var total int64

	query := m.db.WithContext(ctx).Model(&mysql.Comment{})

	if param.PostID != 0 {
		query = query.Where("post_id = ?", param.PostID)
	}

	if err := query.Count(&total).Error; err != nil {
		return mysql.GetPaginatedCommentResponse{}, err
	}

	if param.OrderBy != "" {
		if order := buildOrderClause(param.OrderBy, param.Order, commentOrderableColumns); order != "" {
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
		return mysql.GetPaginatedCommentResponse{}, err
	}

	return mysql.GetPaginatedCommentResponse{Comments: rows, TotalData: int(total)}, nil
}

func (m *MySQL) SaveComment(ctx context.Context, comment mysql.Comment) (mysql.Comment, error) {
	if err := m.db.WithContext(ctx).Save(&comment).Error; err != nil {
		return mysql.Comment{}, err
	}

	return comment, nil
}
