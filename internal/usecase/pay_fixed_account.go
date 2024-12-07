package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type PayFixedAccountUseCase struct {
	repository entity.FixedAccountInterface
}

func NewPayFixedAccountUseCase(repository entity.FixedAccountInterface) *PayFixedAccountUseCase {
	return &PayFixedAccountUseCase{
		repository: repository,
	}
}

func (p *PayFixedAccountUseCase) Execute(fixedAccountID, userID int) *rest_err.RestErr {
	if err := p.repository.PaidFixedAccount(fixedAccountID, userID); err != nil {
		return err
	}

	return nil
}
