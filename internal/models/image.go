package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	URL       string `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	Product   Product
}
