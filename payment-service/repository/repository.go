package repository

import (
	"context"
	"payment-service/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

//go:generate mockery --name IRepository --structname MockRepository --filename mock_db.go --output ./ --outpkg repository
type IRepository interface {
	PayTheBill(ctx context.Context, trans model.Transaction) error
	GetCustomerById(ctx context.Context, id string) (*model.Customer, error)
}

func New(db *gorm.DB) IRepository {
	return &Repository{db: db}
}
