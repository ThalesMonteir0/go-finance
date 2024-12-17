package handlers

import (
	"financial-go/internal/usecase/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("cel")

	deleteUserUseCase := users.NewDeleteUserUseCase(u.UserRepository)

	if err := deleteUserUseCase.Execute(userID); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccess("user deleted successful"))
}
