package repositories

import (
	"freshfinds/internal/models"

	"gorm.io/gorm"
)

type ProductImageRepository struct {
	DB *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{DB: db}
}

func (repo *ProductImageRepository) CreateProductImage(image *models.ProductImage) error {
	if err := repo.DB.Create(image).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductImageRepository) GetProductImagesByProductID(productID uint) ([]models.ProductImage, error) {
	var images []models.ProductImage
	if err := repo.DB.Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
