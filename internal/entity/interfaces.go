package entity

import "financial-go/pkg/rest_err"

type UserRepositoryInterface interface {
	Save(user Users) *rest_err.RestErr
	FindUserByCellphone(cel string) (*Users, *rest_err.RestErr)
	DeleteUserByCellphone(cel string) *rest_err.RestErr
}

type FixedAccountInterface interface {
	Save(account FixedAccounts) *rest_err.RestErr
	FindAllFixedAccountByUserID(userID int) ([]FixedAccounts, *rest_err.RestErr)
	FindAllFixedAccount() ([]FixedAccounts, *rest_err.RestErr)
	DeleteFixedAccount(fixedAccountID, userID int) *rest_err.RestErr
	PaidFixedAccount(fixedAccountID, userID int, paid bool) *rest_err.RestErr
}

type MovementRepositoryInterface interface {
	Save(movement Movements) *rest_err.RestErr
	DeleteMovementByID(movementID, userID int) *rest_err.RestErr
	FindAllMovements(userID int) ([]Movements, *rest_err.RestErr)
}
