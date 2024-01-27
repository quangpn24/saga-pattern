package todo

import (
	"context"
	"warehouse-service/model"
	"warehouse-service/payload"
	"warehouse-service/pkg/constant"

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
