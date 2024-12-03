package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name  string
	Email string
	Cel   string
}

func NewUserEntity(name, email, cel string) *Users {
	return &Users{
		Name:  name,
		Email: email,
		Cel:   cel,
	}
}
