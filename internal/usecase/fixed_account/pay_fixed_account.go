package fixed_account

import (
	"financial-go/pkg/rest_err"
)

func (p *fixedAccountUseCase) PayFixedAccount(fixedAccountID, userID int) *rest_err.RestErr {
	if err := p.repository.PaidFixedAccount(fixedAccountID, userID); err != nil {
		return err
	}

	return nil
}
