package jobs

import "financial-go/internal/usecase/fixed_account"

func UpdateFixedAccountPaidDefault(fixedAccountUseCase fixed_account.FixedAccountUseCaseInterface) {
	fixedAccounts, err := fixedAccountUseCase.FindAll()
	if err != nil {
		//TODO:LOGS'
		return
	}

	for _, fixedAccount := range fixedAccounts {
		if err := fixedAccountUseCase.UpdatePaymentFixedAccount(
			fixedAccount.ID,
			fixedAccount.UserID,
			false); err != nil {
			//TODO:LOGS'
			continue
		}
	}
}
