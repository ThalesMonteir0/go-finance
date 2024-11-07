package entity

type UserRepositoryInterface interface {
	Save(user Users) error
	FindUserByCellphone(cel string) (*Users, error)
	DeleteUserByCellphone(cel string) error
}

type FixedAccountInterface interface {
	Save(account FixedAccounts) error
}
