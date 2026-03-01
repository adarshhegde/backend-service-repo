package api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (handler ApiHandlerImpl) ListAllUsers(c *echo.Context) error {

	err, users := handler.InternalServices.GetUserSvc().ListAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, users)
	return nil
}
