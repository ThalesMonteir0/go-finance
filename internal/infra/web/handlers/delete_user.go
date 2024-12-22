package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("cel")

	if err := u.UserUseCase.Delete(userID); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccess("user deleted successful"))
}
