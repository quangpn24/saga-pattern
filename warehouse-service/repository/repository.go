package repository

import (
	"context"
	"warehouse-service/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

//go:generate mockery --name IRepository --structname MockRepository --filename mock_db.go --output ./ --outpkg repository
type IRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
}

func New(db *gorm.DB) IRepository {
	return &Repository{db: db}
}
