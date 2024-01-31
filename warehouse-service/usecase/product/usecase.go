package product

import (
	"context"
	"warehouse-service/model"
	"warehouse-service/repository"
)

type UseCase struct {
	repo repository.IRepository
}
type IUseCase interface {
	UpdateQuantity(ctx context.Context, products []model.Product) error
}

func NewProductUseCase(repo repository.IRepository) IUseCase {
	return &UseCase{
		repo: repo,
	}
}
