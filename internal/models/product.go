package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	Comments    []Comment `gorm:"foreignKey:ProductID"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
}

type Image struct {
	gorm.Model
	URL       string `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	Product   Product
}

type Like struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
}
