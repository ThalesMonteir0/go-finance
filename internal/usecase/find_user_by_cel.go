package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type FindUserByCelUseCase struct {
	repository entity.UserRepositoryInterface
}

func NewFindUserByCelUseCase(repository entity.UserRepositoryInterface) *FindUserByCelUseCase {
	return &FindUserByCelUseCase{repository: repository}
}

func (f *FindUserByCelUseCase) Execute(cel string) (UserDTOResponse, *rest_err.RestErr) {
	userEntity, err := f.repository.FindUserByCellphone(cel)
	if err != nil {
		return UserDTOResponse{}, err
	}

	userDTO := UserDTOResponse{
		ID:    int(userEntity.ID),
		Name:  userEntity.Name,
		Email: userEntity.Email,
		Cel:   userEntity.Cel,
	}

	return userDTO, nil
}
