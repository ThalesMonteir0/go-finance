package users

import (
	"financial-go/pkg/rest_err"
)

func (f *userUseCase) FindByCellphone(cel string) (UserDTOResponse, *rest_err.RestErr) {
	userEntity, err := f.Repository.FindUserByCellphone(cel)
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
