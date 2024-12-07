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
	if err := m.db.Create(&movement).Error; err != nil {
		return rest_err.NewInternalServeError(err.Error())
	}

	return nil
}

func (m MovementRepository) DeleteMovementByID(movementID, userID int) *rest_err.RestErr {
	if err := m.db.Delete(&entity.Movements{}, "id = ? and user_id = ?", movementID, userID).Error; err != nil {
		return rest_err.NewBadRequestError(err.Error())
	}

	return nil
}

func (m MovementRepository) FindAllMovements(userID int) ([]entity.Movements, *rest_err.RestErr) {
	var movements []entity.Movements

	if err := m.db.Model(&entity.Movements{}).
		Preload("Type").
		Find(&movements, "user_id = ?", userID).
		Error; err != nil {
		return []entity.Movements{}, rest_err.NewBadRequestError(err.Error())
	}

	return movements, nil
}
