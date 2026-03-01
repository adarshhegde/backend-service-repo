package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/adarshhegde/backend-api-repo/internal/cmd/http/dto"
)

func (handler ApiHandlerImpl) CreateUser(c *echo.Context) error {
	var createUserReq dto.CreateUserRequest
	if err := c.Bind(&createUserReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := handler.InternalServices.GetUserSvc().CreateUser(createUserReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
	return nil
}
