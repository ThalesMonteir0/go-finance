package usecase

import "financial-go/internal/entity"

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Cel      string `json:"cel"`
	Password string `json:"password"`
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

func (c *CreateUserUseCase) Execute(user UserDTO) error {
	userEntity := entity.NewUserEntity(user.Name, user.Email, user.Cel, user.Password)

	if err := c.Repository.Save(*userEntity); err != nil {
		return err
	}

	return nil
}
