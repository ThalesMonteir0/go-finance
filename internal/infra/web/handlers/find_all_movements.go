package handlers

import (
	"financial-go/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *MovementHandler) FindAllMovements(c echo.Context) error {
	userID := c.Get("userID").(int)
	findAllMovementsUseCase := usecase.NewFindMovementsUseCase(m.repository)

	movements, err := findAllMovementsUseCase.Execute(userID)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, movements)
}
