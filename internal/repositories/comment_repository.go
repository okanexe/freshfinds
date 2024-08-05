package repositories

import (
	"freshfinds/internal/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (repo *CommentRepository) CreateComment(comment *models.Comment) error {
	if err := repo.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CommentRepository) GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := repo.DB.Where("product_id = ?", productID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
