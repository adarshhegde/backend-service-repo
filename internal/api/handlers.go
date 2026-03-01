package api

import (
	"github.com/labstack/echo/v5"
	"github.com/adarshhegde/backend-api-repo/internal/services"
)

type ApiHandlers interface {
	CreateUser(c *echo.Context) error
	ListAllUsers(c *echo.Context) error
}

type ApiHandlerImpl struct {
	InternalServices services.InternalServices
}

func New(opts services.InternalServicesOpts) ApiHandlers {
	return &ApiHandlerImpl{
		InternalServices: services.New(opts),
	}
}
