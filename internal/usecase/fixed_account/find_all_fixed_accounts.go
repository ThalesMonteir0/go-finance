package fixed_account

import (
	"financial-go/pkg/rest_err"
)

func (f *fixedAccountUseCase) FindAllFixedAccountsByUserID(userID int) ([]FixedAccountDTOResponse, *rest_err.RestErr) {
	var fixedAccountResponse []FixedAccountDTOResponse
	fixedAccounts, err := f.repository.FindAllFixedAccountByUserID(userID)
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

func (f *fixedAccountUseCase) FindAll() ([]FixedAccountDTOResponse, *rest_err.RestErr) {
	var fixedAccountsResponseDTO []FixedAccountDTOResponse
	fixedAccounts, err := f.repository.FindAllFixedAccount()
	if err != nil {
		return nil, err
	}

	for _, fixedAccount := range fixedAccounts {
		fixedAccountResponse := FixedAccountDTOResponse{
			ID:          int(fixedAccount.ID),
			Name:        fixedAccount.Name,
			Desc:        fixedAccount.Desc,
			Value:       fixedAccount.Value,
			PaymentDate: fixedAccount.PaymentDate,
			Paid:        fixedAccount.Paid,
			UserID:      fixedAccount.UserID,
		}

		fixedAccountsResponseDTO = append(fixedAccountsResponseDTO, fixedAccountResponse)
	}

	return fixedAccountsResponseDTO, nil
}
