package usecase

import "financial-go/internal/entity"

type FindUserByCelUseCase struct {
	repository entity.UserRepositoryInterface
}

func NewFindUserByCelUseCase(repository entity.UserRepositoryInterface) *FindUserByCelUseCase {
	return &FindUserByCelUseCase{repository: repository}
}

func (f *FindUserByCelUseCase) Execute(cel string) (UserDTOResponse, error) {
	userEntity, err := f.repository.FindUserByCellphone(cel)
	if err != nil {
		return UserDTOResponse{}, err
	}

	userDTO := UserDTOResponse{
		ID:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
		Cel:   userEntity.Cel,
	}

	return userDTO, nil
}
