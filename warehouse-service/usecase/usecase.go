package usecase

import (
	"warehouse-service/repository"
	"warehouse-service/usecase/product"
)

type UseCase struct {
	ProductUC product.IUseCase
}

func New(repo repository.IRepository) *UseCase {
	return &UseCase{
		ProductUC: product.NewProductUseCase(repo),
	}
}
