package todo

import (
	"warehouse-service/config"
	"warehouse-service/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase  *usecase.UseCase
	validate *validator.Validate
	cfg      *config.Config
}

func Init(group *echo.Group, useCase *usecase.UseCase, validate *validator.Validate, cfg *config.Config) {
	handler := &Handler{
		useCase:  useCase,
		validate: validate,
		cfg:      cfg,
	}

	group.POST("", handler.CreateTodo)
}
