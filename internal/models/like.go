package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
}
