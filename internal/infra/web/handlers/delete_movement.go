package handlers

import (
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (m *MovementHandler) DeleteMovement(c echo.Context) error {
	movementID, err := strconv.Atoi(c.Param("movementID"))
	userID := int(c.Get("userID").(uint))
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid movementID")
		return c.JSON(restErr.Code, NewResponseDataErr(restErr.Message))
	}

	if err := m.movementUseCase.DeleteMovement(movementID, userID); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))

	}

	return c.JSON(http.StatusOK, NewResponseDataSuccess("movement deleted."))
}
