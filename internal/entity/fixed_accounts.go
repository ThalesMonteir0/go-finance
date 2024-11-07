package entity

import "time"

type FixedAccounts struct {
	ID          int
	UserID      int
	Name        string
	Desc        string
	Value       float64
	PaymentDate time.Time
	Paid        bool
}

func NewFixedAccount(name, desc string, value float64, paymentDate time.Time, userID int, paid bool) *FixedAccounts {
	return &FixedAccounts{
		UserID:      userID,
		Name:        name,
		Desc:        desc,
		Value:       value,
		PaymentDate: paymentDate,
		Paid:        paid,
	}
}
