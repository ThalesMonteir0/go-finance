package movement

import (
	"financial-go/pkg/rest_err"
)

func (f *movementUseCase) FindMovements(userID int) ([]MovementDTOResponse, *rest_err.RestErr) {
	var movementsResponse []MovementDTOResponse
	movements, err := f.repository.FindAllMovements(userID)
	if err != nil {
		return nil, err
	}

	for _, movement := range movements {
		movementResponse := MovementDTOResponse{
			ID:        int(movement.ID),
			UserID:    movement.UserID,
			TypeID:    movement.TypeID,
			Value:     movement.Value,
			Desc:      movement.Desc,
			CreatedAT: movement.CreatedAt,
			TypeName:  movement.Type.TypeName,
		}

		movementsResponse = append(movementsResponse, movementResponse)
	}

	return movementsResponse, nil
}
