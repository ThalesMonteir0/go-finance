package usecase

import "financial-go/internal/entity"

type DeleteUser struct {
	repository entity.UserRepositoryInterface
}

func NewDeleteUserUseCase(repository entity.UserRepositoryInterface) *DeleteUser {
	return &DeleteUser{
		repository: repository,
	}
}

func (d *DeleteUser) Execute(cellphone string) error {

	if err := d.repository.DeleteUserByCellphone(cellphone); err != nil {
		return err
	}

	return nil
}
