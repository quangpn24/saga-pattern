package order

import (
	"net/http"
	"order-service/payload"
	"order-service/pkg/appmodel"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req payload.CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
	}
	if err := req.Validate(h.validate); err != nil {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
	}

	if err := h.useCase.OrderUseCase.CreateOrder(c.Request().Context(), req); err != nil {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
	}
	return appmodel.StatusOk(c, "OK")
}
