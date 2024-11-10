package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
	"time"
)

type MovementDTO struct {
	Type  entity.Types `json:"type"`
	Value float64      `json:"value"`
	Desc  string       `json:"desc"`
	Date  time.Time    `json:"date"`
}

type MovementDTOResponse struct {
	ID    int          `json:"id"`
	Type  entity.Types `json:"type"`
	Value float64      `json:"value"`
	Desc  string       `json:"desc"`
	Date  time.Time    `json:"date"`
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
		Type: entity.Types{
			ID: movement.Type.ID,
		},
		Value: movement.Value,
		Desc:  movement.Desc,
		Date:  movement.Date,
	}

	if err := c.repository.Save(movementEntity); err != nil {
		return err
	}

	return nil
}
