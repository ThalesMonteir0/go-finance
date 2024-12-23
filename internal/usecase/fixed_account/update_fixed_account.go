package fixed_account

import (
	"financial-go/pkg/rest_err"
)

func (p *fixedAccountUseCase) UpdatePaymentFixedAccount(fixedAccountID, userID int, paid bool) *rest_err.RestErr {
	if err := p.repository.PaidFixedAccount(fixedAccountID, userID, paid); err != nil {
		return err
	}

	return nil
}
