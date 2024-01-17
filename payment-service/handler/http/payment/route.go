package delivery

import (
	"payment-service/usecase"

	"github.com/labstack/echo/v4"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	_ = &Route{useCase: useCase}

}
