package database

import (
	"financial-go/internal/entity"
	"financial-go/pkg/rest_err"
	"gorm.io/gorm"
)

type MovementRepository struct {
	db *gorm.DB
}

func NewMovementRepository(db *gorm.DB) entity.MovementRepositoryInterface {
	return MovementRepository{
		db: db,
	}
}

func (m MovementRepository) Save(movement entity.Movements) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (m MovementRepository) DeleteMovementByID(movementID int) *rest_err.RestErr {
	//TODO implement me
	panic("implement me")
}

func (m MovementRepository) FindAllMovements() ([]entity.Movements, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
