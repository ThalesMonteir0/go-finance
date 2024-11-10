package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type FindMovementsUseCase struct {
	repository entity.MovementRepositoryInterface
}

func NewFindMovementsUseCase(repository entity.MovementRepositoryInterface) *FindMovementsUseCase {
	return &FindMovementsUseCase{
		repository: repository,
	}
}

func (f *FindMovementsUseCase) Execute() ([]MovementDTOResponse, *rest_err.RestErr) {
	var movementsResponse []MovementDTOResponse
	movements, err := f.repository.FindAllMovements()
	if err != nil {
		return nil, err
	}

	for _, movement := range movements {
		movementResponse := MovementDTOResponse{
			ID:    movement.ID,
			Type:  movement.Type,
			Value: movement.Value,
			Desc:  movement.Desc,
			Date:  movement.Date,
		}

		movementsResponse = append(movementsResponse, movementResponse)
	}

	return movementsResponse, nil
}
