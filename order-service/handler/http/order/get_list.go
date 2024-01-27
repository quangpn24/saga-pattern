package order

import (
	"net/http"
	"order-service/pkg/appmodel"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *OrderHandler) GetListOrders(c echo.Context) error {
	res, err := h.useCase.OrderUseCase.GetListOrders(c.Request().Context())
	if err != nil {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
	}
	return appmodel.StatusOk(c, res)
}
