package usecase

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
)

type FixedAccountDTO struct {
	Name        string  `json:"name"`
	Desc        string  `json:"desc"`
	Value       float64 `json:"value"`
	PaymentDate int     `json:"payment_date"`
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

type CreateFixedAccount struct {
	repository entity.FixedAccountInterface
}

func NewCreateFixedAccount(repository entity.FixedAccountInterface) *CreateFixedAccount {
	return &CreateFixedAccount{
		repository: repository,
	}
}

func (c *CreateFixedAccount) Execute(account FixedAccountDTO) *rest_err.RestErr {
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
