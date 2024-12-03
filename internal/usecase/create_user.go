package usecase

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

type CreateUserUseCase struct {
	Repository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(repository entity.UserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{Repository: repository}
}

func (c *CreateUserUseCase) Execute(user UserDTO) *rest_err.RestErr {
	userEntity := entity.NewUserEntity(user.Name, user.Email, user.Cel)

	if err := c.Repository.Save(*userEntity); err != nil {
		return err
	}

	return nil
}
