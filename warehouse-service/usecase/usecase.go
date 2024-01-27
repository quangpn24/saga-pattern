package usecase

import (
	"warehouse-service/repository"
	"warehouse-service/usecase/todo"
)

type UseCase struct {
	TodoUseCase todo.IUseCase
}

func New(repo repository.IRepository) *UseCase {
	return &UseCase{
		TodoUseCase: todo.NewTodoUseCase(repo),
	}
}
