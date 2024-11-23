package database

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserSaveRepository(db *gorm.DB) entity.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (u UserRepository) Save(user entity.Users) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByCellphone(cel string) (*entity.Users, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) DeleteUserByCellphone(cel string) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}
