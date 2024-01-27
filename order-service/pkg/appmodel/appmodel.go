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
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (err AppError) Error() string {
	return err.Message
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
