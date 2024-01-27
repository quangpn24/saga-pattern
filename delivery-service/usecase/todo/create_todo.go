package todo

import (
	"context"
	"delivery-service/model"
	"delivery-service/payload"
	"delivery-service/pkg/constant"

	"github.com/google/uuid"
)

func (uc *UseCase) CreateTodo(ctx context.Context, req payload.CreateTodoRequest) error {
	todoId := uuid.New().String()
	newTodo := &model.Todo{
		ID:      todoId,
		Status:  constant.TODO_CREATED,
		Content: req.Content,
		Note:    req.Note,
	}
	return uc.repo.CreateTodo(ctx, newTodo)
}
