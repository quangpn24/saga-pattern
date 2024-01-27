package appmodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func StatusOk(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    data,
	})
}
func NewErrorResponse(c echo.Context, code int, err error, message string) error {
	return c.JSON(code, Response{
		code,
		message,
		err,
	})
}
