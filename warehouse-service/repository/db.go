package repository

import (
	"context"
	"warehouse-service/model"
	"warehouse-service/pkg/constant"
)

func (r *Repository) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	var p *model.Product
	err := r.db.WithContext(ctx).Table(constant.ProductTable).Where("id = ?", id).Take(&p).Error
	return p, err
}

func (r *Repository) UpdateProducts(ctx context.Context, products []model.Product) error {
	return r.db.WithContext(ctx).Table(constant.ProductTable).Save(&products).Error
}
