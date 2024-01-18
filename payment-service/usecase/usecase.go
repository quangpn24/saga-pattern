package usecase

import (
	"payment-service/repository"
)

type UseCase struct {
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{}
}
