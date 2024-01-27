package order

import (
	"errors"
	"net/http"
	"order-service/pkg/appmodel"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (h *OrderHandler) GetOrderById(c echo.Context) error {
	orderId := c.Param("id")
	if orderId == "" {
		err := errors.New("order id not null")
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
	}

	order, err := h.useCase.OrderUseCase.GetById(c.Request().Context(), orderId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
	}
	if err == gorm.ErrRecordNotFound {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
	}

	return appmodel.StatusOk(c, order)
}
