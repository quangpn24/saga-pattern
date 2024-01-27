package order

import (
	"context"
	"order-service/kafka"
	"order-service/model"
	"order-service/payload"
	"order-service/repository"
)

type UseCase struct {
	repo     repository.IRepository
	producer kafka.IProducer
}
type IUseCase interface {
	GetListOrders(ctx context.Context) ([]model.Order, error)
	GetById(ctx context.Context, id string) (*model.Order, error)
	CreateOrder(ctx context.Context, req payload.CreateOrderRequest) error
	CancelOrder(ctx context.Context, id string) error
	RejectOrder(ctx context.Context, id string, note string) error
	UpdateStatus(ctx context.Context, id string, status string) error
}

func NewOrderUseCase(repo repository.IRepository) IUseCase {
	return &UseCase{
		repo:     repo,
		producer: kafka.NewProducer(),
	}
}
