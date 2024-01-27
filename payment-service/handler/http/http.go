package http

import (
	"net/http"
	"payment-service/config"
	"payment-service/handler/http/customer"
	"payment-service/handler/http/healthcheck"
	auth "payment-service/pkg/appmiddleware"
	"payment-service/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPHandler(uc *usecase.UseCase, cfg *config.Config) *echo.Echo {
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

	//middleware
	skipperPath := []string{
		"/health-check",
	}
	e.Use(auth.NewAuthentication("header:Authorization", "Bearer", skipperPath).Middleware())

	//validate
	validate := validator.New()

	// Health check use for microservice
	healthcheck.Init(e.Group("/health-check"))

	// APIs
	api := e.Group("/api")

	// customer route
	customer.Init(api.Group("/customers"), uc, validate, cfg)
	return e
}
