package users

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cel   string `json:"cel"`
}

type UserDTOResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cel   string `json:"cel"`
}

type userUseCase struct {
	Repository entity.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	Create(user UserDTO) *rest_err.RestErr
	Delete(cellphone string) *rest_err.RestErr
	FindByCellphone(cell string) (UserDTOResponse, *rest_err.RestErr)
}

func NewUserUseCase(repository entity.UserRepositoryInterface) UserUseCaseInterface {
	return &userUseCase{Repository: repository}
}

func (c *userUseCase) Create(user UserDTO) *rest_err.RestErr {
	userEntity := entity.NewUserEntity(user.Name, user.Email, user.Cel)

	if err := c.Repository.Save(*userEntity); err != nil {
		return err
	}

	return nil
}
