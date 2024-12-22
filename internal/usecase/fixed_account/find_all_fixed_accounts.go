package fixed_account

import (
	"financial-go/pkg/rest_err"
)

func (f *fixedAccountUseCase) FindAllFixedAccounts(userID int) ([]FixedAccountDTOResponse, *rest_err.RestErr) {
	var fixedAccountResponse []FixedAccountDTOResponse
	fixedAccounts, err := f.repository.FindAllFixedAccount(userID)
	if err != nil {
		return []FixedAccountDTOResponse{}, err
	}

	for _, fixedAccount := range fixedAccounts {
		fixedAccountDTO := FixedAccountDTOResponse{
			ID:          int(fixedAccount.ID),
			Name:        fixedAccount.Name,
			Desc:        fixedAccount.Desc,
			Value:       fixedAccount.Value,
			PaymentDate: fixedAccount.PaymentDate,
			Paid:        fixedAccount.Paid,
			UserID:      fixedAccount.UserID,
		}
		fixedAccountResponse = append(fixedAccountResponse, fixedAccountDTO)
	}

	return fixedAccountResponse, nil
}