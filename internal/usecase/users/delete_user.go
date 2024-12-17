package users

import (
	"financial-go/pkg/rest_err"
)

func (d *userUseCase) Delete(cellphone string) *rest_err.RestErr {
	if err := d.Repository.DeleteUserByCellphone(cellphone); err != nil {
		return err
	}

	return nil
}
