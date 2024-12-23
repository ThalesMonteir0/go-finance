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
	if err := f.db.Create(&account).Error; err != nil {
		return rest_err.NewBadRequestError(err.Error())
	}

	return nil
}

func (f FixedAccountRepository) FindAllFixedAccountByUserID(userID int) ([]entity.FixedAccounts, *rest_err.RestErr) {
	var accounts []entity.FixedAccounts

	if err := f.db.Find(&accounts, "user_id = ?", userID).Error; err != nil {
		return []entity.FixedAccounts{}, rest_err.NewBadRequestError(err.Error())
	}

	return accounts, nil
}

func (f FixedAccountRepository) DeleteFixedAccount(fixedAccountID, userID int) *rest_err.RestErr {
	if err := f.db.Delete(&entity.FixedAccounts{}, "id = ? and user_id = ?", fixedAccountID, userID).Error; err != nil {
		return rest_err.NewBadRequestError(err.Error())
	}

	return nil
}

func (f FixedAccountRepository) PaidFixedAccount(fixedAccountID, userID int, paid bool) *rest_err.RestErr {
	if err := f.db.Model(&entity.FixedAccounts{}).
		Where("id = ? and user_id = ?", fixedAccountID, userID).
		Update("paid", paid).
		Error; err != nil {
		return rest_err.NewBadRequestError(err.Error())
	}

	return nil
}

func (f FixedAccountRepository) FindAllFixedAccount() ([]entity.FixedAccounts, *rest_err.RestErr) {
	var accounts []entity.FixedAccounts

	if err := f.db.Find(&accounts).Error; err != nil {
		return []entity.FixedAccounts{}, rest_err.NewBadRequestError(err.Error())
	}

	return accounts, nil
}
