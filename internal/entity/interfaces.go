package entity

import "financial-go/pkg/rest_err"

type UserRepositoryInterface interface {
	Save(user Users) *rest_err.RestErr
	FindUserByCellphone(cel string) (*Users, *rest_err.RestErr)
	DeleteUserByCellphone(cel string) *rest_err.RestErr
}

type FixedAccountInterface interface {
	Save(account FixedAccounts) *rest_err.RestErr
}
