package repository

import (
	"context"
	"order-service/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

//go:generate mockery --name IRepository --structname MockRepository --filename mock_db.go --output ./ --outpkg repository
type IRepository interface {
	GetListOrders(ctx context.Context) ([]model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	GetById(ctx context.Context, id string) (*model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) error
	UpdateStatus(ctx context.Context, id string, status string, note string) error
}

func New(db *gorm.DB) IRepository {
	return &Repository{db: db}
}
