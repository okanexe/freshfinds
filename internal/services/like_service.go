package services

import (
	"freshfinds/internal/models"
	"freshfinds/internal/repositories"
)

type LikeService struct {
	LikeRepo *repositories.LikeRepository
}

func NewLikeService(likeRepo *repositories.LikeRepository) *LikeService {
	return &LikeService{LikeRepo: likeRepo}
}

func (s *LikeService) AddLike(userID, productID uint) error {
	return s.LikeRepo.AddLike(&models.Like{
		UserID:    userID,
		ProductID: productID,
	})
}

func (s *LikeService) RemoveLike(userID, productID uint) error {
	return s.LikeRepo.RemoveLike(userID, productID)
}

func (s *LikeService) GetLikesByProductID(productID uint) ([]models.Like, error) {
	return s.LikeRepo.GetLikesByProductID(productID)
}

func (s *LikeService) GetUserLikes(userID uint) ([]models.Like, error) {
	return s.LikeRepo.GetUserLikes(userID)
}
