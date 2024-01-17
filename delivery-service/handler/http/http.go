package http

import (
	"net/http"
	"order-service/handler/http/healthcheck"
	"order-service/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e = echo.New()
	)

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check use for microservice
	healthcheck.Init(e.Group("/health-check"))

	// APIs
	//api := e.Group("/api")

	return e
}