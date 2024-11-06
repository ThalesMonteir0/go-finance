package entity

type UserRepositoryInterface interface {
	Save(user Users) error
	FindUserByCellphone(cel string) (*Users, error)
}
