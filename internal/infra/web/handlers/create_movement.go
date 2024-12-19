package handlers

import (
	"financial-go/internal/usecase/movement"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MovementHandler struct {
	movementUseCase movement.MovementUseCaseInterface
}

func NewMovementHandler(useCase movement.MovementUseCaseInterface) MovementHandler {
	return MovementHandler{
		movementUseCase: useCase,
	}
}

func (m *MovementHandler) CreateMovement(c echo.Context) error {
	var movementDTO movement.MovementDTO
	movementDTO.UserID = int(c.Get("userID").(uint))

	if err := c.Bind(&movementDTO); err != nil {
		restErr := rest_err.NewBadRequestError("invalid body")
		return c.JSON(restErr.Code, NewResponseDataErr(restErr.Message))
	}

	if err := m.movementUseCase.Create(movementDTO); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("movement created"))
}
