package todo

import (
	"context"
	"errors"
	"testing"
	"warehouse-service/model"
	"warehouse-service/payload"
	"warehouse-service/pkg/constant"
	"warehouse-service/repository"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUseCase_CreateTodo(t *testing.T) {
	mockRepo := repository.NewMockRepository(t)
	uc := NewTodoUseCase(mockRepo)

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {

		req := payload.CreateTodoRequest{
			Content: "Write unit test",
			Note:    "",
		}

		newTodo := model.Todo{
			ID:      uuid.New().String(),
			Status:  constant.TODO_CREATED,
			Content: req.Content,
			Note:    req.Note,
		}

		mockRepo.On("CreateTodo", ctx, mock.MatchedBy(func(todo *model.Todo) bool {
			return cmp.Equal(newTodo, *todo,
				cmpopts.IgnoreFields(model.Todo{}, "ID"),
				cmpopts.IgnoreFields(model.Todo{}, "CreatedAt"),
				cmpopts.IgnoreFields(model.Todo{}, "UpdatedAt"))
		})).Return(nil).Once()
		err := uc.CreateTodo(ctx, req)

		require.NoError(t, err)
	})
	t.Run("unexpected error", func(t *testing.T) {

		req := payload.CreateTodoRequest{
			Content: "Write unit test",
			Note:    "",
		}

		newTodo := model.Todo{
			ID:      uuid.New().String(),
			Status:  constant.TODO_CREATED,
			Content: req.Content,
			Note:    req.Note,
		}

		mockRepo.On("CreateTodo", ctx, mock.MatchedBy(func(todo *model.Todo) bool {
			return cmp.Equal(newTodo, *todo,
				cmpopts.IgnoreFields(model.Todo{}, "ID"),
				cmpopts.IgnoreFields(model.Todo{}, "CreatedAt"),
				cmpopts.IgnoreFields(model.Todo{}, "UpdatedAt"))
		})).Return(errors.New("unexpected error")).Once()
		err := uc.CreateTodo(ctx, req)

		require.Error(t, err)
	})
}
