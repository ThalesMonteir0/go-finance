package fixed_account

import (
	"financial-go/pkg/rest_err"
)

func (d *fixedAccountUseCase) DeleteFixedAccount(fixedAccountID, userID int) *rest_err.RestErr {
	if err := d.repository.DeleteFixedAccount(fixedAccountID, userID); err != nil {
		return err
	}

	return nil
}
