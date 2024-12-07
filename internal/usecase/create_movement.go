package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
	"time"
)

type MovementDTO struct {
	TypeID int     `json:"type_id"`
	UserID int     `json:"user_id"`
	Value  float64 `json:"value"`
	Desc   string  `json:"desc"`
}

type MovementDTOResponse struct {
	ID        int       `json:"id"`
	TypeID    int       `json:"type_id"`
	UserID    int       `json:"user_id"`
	Value     float64   `json:"value"`
	Desc      string    `json:"desc"`
	TypeName  string    `json:"type_name"`
	CreatedAT time.Time `json:"created_at"`
}

type CreateMovementUseCase struct {
	repository entity.MovementRepositoryInterface
}

func NewCreateMovementUseCase(repository entity.MovementRepositoryInterface) *CreateMovementUseCase {
	return &CreateMovementUseCase{
		repository: repository,
	}
}

func (c *CreateMovementUseCase) Execute(movement MovementDTO) *rest_err.RestErr {
	movementEntity := entity.Movements{
		TypeID: movement.TypeID,
		Value:  movement.Value,
		Desc:   movement.Desc,
		UserID: movement.UserID,
	}

	if err := c.repository.Save(movementEntity); err != nil {
		return err
	}

	return nil
}
