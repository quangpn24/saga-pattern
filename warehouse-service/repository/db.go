package repository

import (
	"context"
	"warehouse-service/model"
	"warehouse-service/pkg/constant"
)

func (r *Repository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	err := r.db.WithContext(ctx).Table(constant.TodoTable).Create(&todo).Error
	return err
}
