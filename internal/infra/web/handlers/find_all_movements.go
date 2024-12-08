package handlers

import (
	"financial-go/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *MovementHandler) FindAllMovements(c echo.Context) error {
	userID := int(c.Get("userID").(uint))
	findAllMovementsUseCase := usecase.NewFindMovementsUseCase(m.repository)

	movements, err := findAllMovementsUseCase.Execute(userID)
	if err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccessWithData(movements, "find all movements with success"))
}
