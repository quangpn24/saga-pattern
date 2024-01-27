package usecase

import (
	"payment-service/repository"
	"payment-service/usecase/customer"
)

type UseCase struct {
	CustomerUseCase customer.IUseCase
}

func New(repo repository.IRepository) *UseCase {
	return &UseCase{
		CustomerUseCase: customer.NewCustomerUseCase(repo),
	}
}
