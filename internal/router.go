package internal

import (
	"freshfinds/internal/handlers"
	"freshfinds/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handlers.AuthHandler, productHandler *handlers.ProductHandler, likeHandler *handlers.LikeHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	products := r.Group("/api/v1/products")
	{
		products.GET("/", productHandler.GetAllProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.Use(middleware.AuthMiddleware)
		{
			products.POST("/", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			products.POST("/:productID/comments", productHandler.AddComment)
			products.GET("/:productID/comments", productHandler.GetCommentsByProductID)
			products.POST("/:productID/images", productHandler.AddProductImage)
			products.GET("/:productID/images", productHandler.GetProductImagesByProductID)
			products.POST("/:productID/like", likeHandler.AddLike)
			products.DELETE("/:productID/like", likeHandler.RemoveLike)
			products.GET("/:productID/likes", likeHandler.GetLikesByProductID)
		}
	}

	r.GET("/api/v1/user/likes", likeHandler.GetUserLikes)

	return r
}
