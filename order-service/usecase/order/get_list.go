package order

import (
	"context"
	"order-service/model"
)

func (uc *UseCase) GetListOrders(ctx context.Context) ([]model.Order, error) {
	return uc.repo.GetListOrders(ctx)
}
