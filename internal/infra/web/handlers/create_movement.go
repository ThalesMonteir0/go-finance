package handlers

import (
	"financial-go/internal/entity"
	"financial-go/internal/usecase"
	"financial-go/pkg/rest_err"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MovementHandler struct {
	repository entity.MovementRepositoryInterface
}

func NewMovementHandler(repo entity.MovementRepositoryInterface) MovementHandler {
	return MovementHandler{
		repository: repo,
	}
}

func (m *MovementHandler) CreateMovement(c echo.Context) error {
	var movementDTO usecase.MovementDTO
	movementDTO.UserID = int(c.Get("userID").(uint))

	if err := c.Bind(&movementDTO); err != nil {
		restErr := rest_err.NewBadRequestError("invalid body")
		return c.JSON(restErr.Code, NewResponseDataErr(restErr.Message))
	}

	createMovement := usecase.NewCreateMovementUseCase(m.repository)

	if err := createMovement.Execute(movementDTO); err != nil {
		return c.JSON(err.Code, NewResponseDataErr(err.Message))
	}

	return c.JSON(http.StatusCreated, NewResponseDataSuccess("movement created"))
}
