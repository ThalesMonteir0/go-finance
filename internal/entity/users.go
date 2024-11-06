package entity

type Users struct {
	ID       int
	Name     string
	Email    string
	Cel      string
	Password string
}

func NewUserEntity(name, email, cel, password string) *Users {
	return &Users{
		Name:     name,
		Email:    email,
		Cel:      cel,
		Password: password,
	}
}

func (u *Users) PasswordHash() {

}
