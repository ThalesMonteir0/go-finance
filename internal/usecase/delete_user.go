package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type DeleteUser struct {
	repository entity.UserRepositoryInterface
}

func NewDeleteUserUseCase(repository entity.UserRepositoryInterface) *DeleteUser {
	return &DeleteUser{
		repository: repository,
	}
}

func (d *DeleteUser) Execute(cellphone string) *rest_err.RestErr {

	if err := d.repository.DeleteUserByCellphone(cellphone); err != nil {
		return err
	}

	return nil
}
