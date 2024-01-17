package order

import (
	"go-base/usecase"

	"github.com/labstack/echo/v4"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

}
