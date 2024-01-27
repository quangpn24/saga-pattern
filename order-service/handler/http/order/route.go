package order

import (
	"order-service/config"
	"order-service/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	useCase  *usecase.UseCase
	validate *validator.Validate
	cfg      *config.Config
}

func Init(group *echo.Group, useCase *usecase.UseCase, validate *validator.Validate, cfg *config.Config) {
	handler := &OrderHandler{
		useCase:  useCase,
		validate: validate,
		cfg:      cfg,
	}

	group.GET("", handler.GetListOrders)
	group.POST("", handler.CreateOrder)
	group.GET("/:id", handler.GetOrderById)
	group.PUT("/cancel/:id", handler.CancelOrder)
}
