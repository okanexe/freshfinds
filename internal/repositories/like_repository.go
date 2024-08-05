package repositories

import (
	"freshfinds/internal/models"

	"gorm.io/gorm"
)

type LikeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{DB: db}
}

func (repo *LikeRepository) AddLike(like *models.Like) error {
	if err := repo.DB.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (repo *LikeRepository) RemoveLike(userID, productID uint) error {
	if err := repo.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Like{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *LikeRepository) GetLikesByProductID(productID uint) ([]models.Like, error) {
	var likes []models.Like
	if err := repo.DB.Where("product_id = ?", productID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (repo *LikeRepository) GetUserLikes(userID uint) ([]models.Like, error) {
	var likes []models.Like
	if err := repo.DB.Where("user_id = ?", userID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}
