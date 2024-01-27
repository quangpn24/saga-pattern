package repository

import (
	"context"
	"order-service/model"
	"order-service/pkg/constant"
)

func (r *Repository) GetListOrders(ctx context.Context) ([]model.Order, error) {
	var res []model.Order
	err := r.db.WithContext(ctx).Table(constant.OrderTable).Find(&res).Error
	return res, err
}
func (r *Repository) CreateOrder(ctx context.Context, order *model.Order) error {
	err := r.db.WithContext(ctx).Table(constant.OrderTable).Create(&order).Error
	return err
}
func (r *Repository) GetById(ctx context.Context, id string) (*model.Order, error) {
	var order *model.Order
	err := r.db.WithContext(ctx).Table(constant.OrderTable).Where("id = ?", id).Take(&order).Error
	return order, err
}
func (r *Repository) UpdateOrder(ctx context.Context, order *model.Order) error {
	err := r.db.WithContext(ctx).Table(constant.OrderTable).Save(&order).Error
	return err
}
func (r *Repository) UpdateStatus(ctx context.Context, id string, status string, note string) error {
	err := r.db.WithContext(ctx).Table(constant.OrderTable).Where("id = ?", id).Updates(map[string]interface{}{"status": status, "note": note}).Error
	return err
}
