package repositories

import (
	"errors"

	"freshfinds/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	if err := repo.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := repo.DB.Preload("Comments").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	if err := repo.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) DeleteProduct(id uint) error {
	if err := repo.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := repo.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
