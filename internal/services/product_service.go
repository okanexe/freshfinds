package services

import (
	"errors"

	"freshfinds/internal/models"
	"freshfinds/internal/repositories"
)

type ProductService struct {
	ProductRepo      *repositories.ProductRepository
	CommentRepo      *repositories.CommentRepository
	ProductImageRepo *repositories.ProductImageRepository
}

func NewProductService(productRepo *repositories.ProductRepository, commentRepo *repositories.CommentRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo, CommentRepo: commentRepo}
}

func (s *ProductService) CreateProduct(name, description string, price float64, userID uint) (*models.Product, error) {
	product := &models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		UserID:      userID,
	}

	if err := s.ProductRepo.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.ProductRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s *ProductService) UpdateProduct(id uint, name, description string, price float64) (*models.Product, error) {
	product, err := s.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price

	if err := s.ProductRepo.UpdateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if id == 0 {
		return errors.New("id must be valid")
	}
	return s.ProductRepo.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.ProductRepo.GetAllProducts()
}

func (s *ProductService) AddComment(content string, userID, productID uint) (*models.Comment, error) {
	comment := &models.Comment{
		Content:   content,
		UserID:    userID,
		ProductID: productID,
	}

	if err := s.CommentRepo.CreateComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *ProductService) GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	return s.CommentRepo.GetCommentsByProductID(productID)
}

func (s *ProductService) AddProductImage(url string, productID uint) (*models.Image, error) {
	image := &models.Image{
		URL:       url,
		ProductID: productID,
	}

	if err := s.ProductImageRepo.CreateProductImage(image); err != nil {
		return nil, err
	}

	return image, nil
}

func (s *ProductService) GetProductImagesByProductID(productID uint) ([]models.Image, error) {
	return s.ProductImageRepo.GetProductImagesByProductID(productID)
}
