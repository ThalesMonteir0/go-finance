package movement

import (
	"financial-go/pkg/rest_err"
)

func (d *movementUseCase) DeleteMovement(movementID, userID int) *rest_err.RestErr {
	if err := d.repository.DeleteMovementByID(movementID, userID); err != nil {
		return err
	}

	return nil
}
