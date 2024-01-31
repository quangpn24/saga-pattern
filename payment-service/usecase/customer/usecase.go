package customer

import (
	"context"
	kafka2 "payment-service/kafka"
	"payment-service/repository"
)

type UseCase struct {
	repo repository.IRepository
}
type IUseCase interface {
	PayTheBill(ctx context.Context, req kafka2.OrderCreatedMessage) (string, error)
	Refund(ctx context.Context, transactionId string) error
}

func NewCustomerUseCase(repo repository.IRepository) IUseCase {
	return &UseCase{
		repo: repo,
	}
}
