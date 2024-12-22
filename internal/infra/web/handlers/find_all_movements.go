package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *MovementHandler) FindAllMovements(c echo.Context) error {
	userID := int(c.Get("userID").(uint))
	
	movements, err := m.movementUseCase.FindMovements(userID)
	if err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusOK, NewResponseDataSuccessWithData(movements, "find all movements with success"))
}
