package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type DeleteFixedAccountUseCase struct {
	repository entity.FixedAccountInterface
}

func NewDeleteFixedAccountUseCase(repository entity.FixedAccountInterface) *DeleteFixedAccountUseCase {
	return &DeleteFixedAccountUseCase{
		repository: repository,
	}
}

func (d *DeleteFixedAccountUseCase) Execute(fixedAccountID, userID int) *rest_err.RestErr {
	if err := d.repository.DeleteFixedAccount(fixedAccountID, userID); err != nil {
		return err
	}

	return nil
}
