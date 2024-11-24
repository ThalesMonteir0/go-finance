package handlers

import (
	"financial-go/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("cel")

	deleteUserUseCase := usecase.NewDeleteUserUseCase(u.userRepository)

	if err := deleteUserUseCase.Execute(userID); err != nil {
		return c.JSON(err.Code, echo.Map{"error": err.Message})
	}

	return c.JSON(http.StatusOK, "user deleted successful")
}
