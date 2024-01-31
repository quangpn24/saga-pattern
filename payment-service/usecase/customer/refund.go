package customer

import "context"

func (uc *UseCase) Refund(ctx context.Context, transactionId string) error {
	return uc.repo.Refund(ctx, transactionId)
}
