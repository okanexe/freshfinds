package main

import (
	"freshfinds/config"
	"freshfinds/internal"
	"freshfinds/internal/handlers"
	"freshfinds/internal/models"
	"freshfinds/internal/repositories"
	"freshfinds/internal/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()

	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Comment{})
	if err != nil {
		panic("failed to auto migrate")
	}

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	likeRepo := repositories.NewLikeRepository(db)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	productService := services.NewProductService(productRepo, commentRepo)
	likeService := services.NewLikeService(likeRepo)

	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	likeHandler := handlers.NewLikeHandler(likeService)

	r := internal.SetupRouter(authHandler, productHandler, likeHandler)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
