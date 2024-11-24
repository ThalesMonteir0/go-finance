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
	if err := u.db.Create(&user).Error; err != nil {
		return rest_err.NewInternalServeError(err.Error())
	}

	return nil
}

func (u UserRepository) FindUserByCellphone(cel string) (*entity.Users, *rest_err.RestErr) {
	var user entity.Users
	if err := u.db.First(&user, "cel = ?", cel).Error; err != nil {
		return &entity.Users{}, rest_err.NewBadRequestError(err.Error())
	}

	return &user, nil
}

func (u UserRepository) DeleteUserByCellphone(cel string) *rest_err.RestErr {
	if err := u.db.Delete(&entity.Users{}, "cel =?", cel).Error; err != nil {
		return rest_err.NewBadRequestError(err.Error())
	}

	return nil
}
