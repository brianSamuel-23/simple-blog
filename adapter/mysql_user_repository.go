package adapter

import (
	"context"
	"errors"

	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"simple-blog.com/entity"
	"simple-blog.com/impl/repo/mysql/mysql"
	"simple-blog.com/user/user"
)

const mysqlErrDuplicateEntry = 1062

type UserRepository interface {
	GetOneUser(ctx context.Context, param mysql.GetOneUserParam) (mysql.User, error)
	CreateUser(ctx context.Context, row mysql.User) (mysql.User, error)
}

type mysqlUserRepository struct {
	repo UserRepository
}

func NewMysqlUserRepository(repo UserRepository) *mysqlUserRepository {
	return &mysqlUserRepository{repo}
}

func (r *mysqlUserRepository) GetOne(ctx context.Context, param user.GetOneUserParam) (entity.User, error) {
	row, err := r.repo.GetOneUser(ctx, mysql.GetOneUserParam{ID: param.ID, Name: param.Name, Email: param.Email})

	if err != nil && errors.Is(gorm.ErrRecordNotFound, err) {
		return entity.User{}, user.ErrUserNotFound
	} else if err != nil {
		return entity.User{}, err
	}

	return toEntityUser(row), nil
}

func (r *mysqlUserRepository) Create(ctx context.Context, usr entity.User) (entity.User, error) {
	row, err := r.repo.CreateUser(ctx, mysql.User{
		Name:         usr.Name,
		Email:        usr.Email,
		PasswordHash: usr.PasswordHash,
	})

	var mysqlErr *sqldriver.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == mysqlErrDuplicateEntry {
		return entity.User{}, user.ErrEmailAlreadyExists
	} else if err != nil {
		return entity.User{}, err
	}

	return toEntityUser(row), nil
}

func toEntityUser(m mysql.User) entity.User {
	return entity.User{
		ID:           m.ID,
		Name:         m.Name,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Timestamps: entity.Timestamps{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}
