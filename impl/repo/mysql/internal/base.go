package internal

import (
	"strings"

	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func NewMySQL(db *gorm.DB) *MySQL {
	return &MySQL{db}
}

// buildOrderClause validates orderBy against allowedColumns and direction
// against ASC/DESC, returning "" if either is invalid. This prevents SQL
// injection through GORM's Order(), which does not parameterize its input.
func buildOrderClause(orderBy, direction string, allowedColumns map[string]bool) string {
	if !allowedColumns[strings.ToLower(orderBy)] {
		return ""
	}

	order := orderBy
	switch strings.ToUpper(direction) {
	case "ASC":
		order += " ASC"
	case "DESC":
		order += " DESC"
	}

	return order
}
