package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type DeleteMovementUseCase struct {
	repository entity.MovementRepositoryInterface
}

func NewDeleteMovementUseCase(repository entity.MovementRepositoryInterface) *DeleteMovementUseCase {
	return &DeleteMovementUseCase{
		repository: repository,
	}
}

func (d *DeleteMovementUseCase) Execute(movementID, userID int) *rest_err.RestErr {
	if err := d.repository.DeleteMovementByID(movementID, userID); err != nil {
		return err
	}

	return nil
}
