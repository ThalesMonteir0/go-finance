package entity

import "time"

type FixedAccounts struct {
	ID          int
	Name        string
	Desc        string
	Value       float64
	PaymentDate time.Time
	UserID      int
}
