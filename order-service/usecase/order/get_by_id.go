package order

import (
	"context"
	"order-service/model"
)

func (uc *UseCase) GetById(ctx context.Context, id string) (*model.Order, error) {
	return uc.repo.GetById(ctx, id)
}
