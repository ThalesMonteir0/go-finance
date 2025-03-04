package fixed_account

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type FixedAccountDTO struct {
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Value       float64 `json:"value"`
	PaymentDate int     `json:"paymentDay"`
	Paid        bool    `json:"paid"`
	UserID      int     `json:"user_id"`
}

type FixedAccountDTOResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Value       float64 `json:"value"`
	PaymentDate int     `json:"payment_date"`
	Paid        bool    `json:"paid"`
	UserID      int     `json:"user_id"`
}

type fixedAccountUseCase struct {
	repository entity.FixedAccountInterface
}

type FixedAccountUseCaseInterface interface {
	CreateFixedAccount(FixedAccountDTO) *rest_err.RestErr
	DeleteFixedAccount(fixedAccountID, userID int) *rest_err.RestErr
	FindAllFixedAccountsByUserID(userID int) ([]FixedAccountDTOResponse, *rest_err.RestErr)
	UpdatePaymentFixedAccount(fixedAccountID, userID int, paid bool) *rest_err.RestErr
	FindAll() ([]FixedAccountDTOResponse, *rest_err.RestErr)
}

func NewFixedAccount(repository entity.FixedAccountInterface) FixedAccountUseCaseInterface {
	return &fixedAccountUseCase{
		repository: repository,
	}
}

func (c *fixedAccountUseCase) CreateFixedAccount(account FixedAccountDTO) *rest_err.RestErr {
	fixedAccountEntity := entity.NewFixedAccount(
		account.Name,
		account.Desc,
		account.Value,
		account.PaymentDate,
		account.UserID,
		account.Paid,
	)

	if err := c.repository.Save(*fixedAccountEntity); err != nil {
		return err
	}

	return nil
}
