package entity

import (
	"gorm.io/gorm"
)

type Movements struct {
	gorm.Model
	TypeID int
	UserID int
	Value  float64
	Desc   string
	Type   Types `gorm:"foreignKey:TypeID"`
}
