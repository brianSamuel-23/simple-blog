package internal

import (
	"context"

	"simple-blog.com/impl/repo/mysql/mysql"
)

func (m *MySQL) GetOneUser(ctx context.Context, param mysql.GetOneUserParam) (mysql.User, error) {
	var row mysql.User

	query := m.db.WithContext(ctx)
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Name != "" {
		query = query.Where("name = ?", param.Name)
	}
	if param.Email != "" {
		query = query.Where("email = ?", param.Email)
	}

	if err := query.First(&row).Error; err != nil {
		return mysql.User{}, err
	}

	return row, nil
}

func (m *MySQL) CreateUser(ctx context.Context, row mysql.User) (mysql.User, error) {
	if err := m.db.WithContext(ctx).Create(&row).Error; err != nil {
		return mysql.User{}, err
	}

	return row, nil
}
