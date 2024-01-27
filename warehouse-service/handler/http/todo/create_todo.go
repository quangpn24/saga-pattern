package todo

import (
	"net/http"
	"warehouse-service/payload"
	"warehouse-service/pkg/appmodel"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateTodo(c echo.Context) error {
	var req payload.CreateTodoRequest

	//parse request
	if err := c.Bind(&req); err != nil {
		logrus.Error("Invalid input: " + err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, "Invalid input")
	}

	//validate
	if err := req.Validate(h.validate); err != nil {
		logrus.Error("Validate fail: " + err.Error())
		return appmodel.NewErrorResponse(c, http.StatusBadRequest, err, "Invalid input")
	}

	//create
	if err := h.useCase.TodoUseCase.CreateTodo(c.Request().Context(), req); err != nil {
		logrus.Error(err.Error())
		return appmodel.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
	}

	return appmodel.StatusOk(c, "Todo created")
}
