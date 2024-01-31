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
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	UpdateProducts(ctx context.Context, products []model.Product) error
}

func New(db *gorm.DB) IRepository {
	return &Repository{db: db}
}
