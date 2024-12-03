package entity

import (
	"gorm.io/gorm"
	"time"
)

type FixedAccounts struct {
	gorm.Model
	UserID      int
	Name        string
	Desc        string `gorm:"column:description"`
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
