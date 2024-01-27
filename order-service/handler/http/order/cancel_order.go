package order

import (
	"errors"
	"net/http"
	"order-service/pkg/appmodel"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *OrderHandler) CancelOrder(c echo.Context) error {
	orderId := c.Param("id")
	if orderId == "" {
		err := errors.New("order id not null")
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
	}

	if err := h.useCase.OrderUseCase.CancelOrder(c.Request().Context(), orderId); err != nil {
		appErr, _ := err.(appmodel.AppError)
		logrus.Error(appErr.Error())
		return appmodel.NewErrorResponse(c, appErr.Code, appErr, appErr.Message)
	}

	return appmodel.StatusOk(c, "Cancel successfully")
}
