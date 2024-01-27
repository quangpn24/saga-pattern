package order

import (
	"context"
	"order-service/pkg/constant"
)

func (uc *UseCase) RejectOrder(ctx context.Context, id string, note string) error {
	return uc.repo.UpdateStatus(ctx, id, string(constant.ORDER_REJECTED), note)
}
