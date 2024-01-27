package usecase

import (
	"order-service/repository"
	"order-service/usecase/order"
)

type UseCase struct {
	OrderUseCase order.IUseCase
}

func New(repo repository.IRepository) *UseCase {
	return &UseCase{
		OrderUseCase: order.NewOrderUseCase(repo),
	}
}
