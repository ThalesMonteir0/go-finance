package handlers

import (
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *MovementHandler) DeleteMovement(c echo.Context) error {
	movementID, err := strconv.Atoi(c.Param("movementID"))
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid movementID")
		return c.JSON(restErr.Code, restErr.Message)
	}

	deleteUseCase := usecase.NewDeleteMovementUseCase(m.repository)

	if err := deleteUseCase.Execute(movementID); err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(http.StatusOK, "movement deleted.")
}
