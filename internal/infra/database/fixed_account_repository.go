package database

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
	"gorm.io/gorm"
)

type FixedAccountRepository struct {
	db *gorm.DB
}

func NewFixedAccountRepository(db *gorm.DB) entity.FixedAccountInterface {
	return &FixedAccountRepository{
		db: db,
	}
}

func (f FixedAccountRepository) Save(account entity.FixedAccounts) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (f FixedAccountRepository) FindAllFixedAccount(userID int) ([]entity.FixedAccounts, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}

func (f FixedAccountRepository) DeleteFixedAccount(fixedAccountID int) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (f FixedAccountRepository) PaidFixedAccount(fixedAccountID int) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}
