package routes

import (
	"github.com/erkanattt/go-rest-crud-project/internal/delivery"
	"github.com/erkanattt/go-rest-crud-project/internal/repository"
	"github.com/erkanattt/go-rest-crud-project/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := delivery.NewProductHandler(productService)

	products := r.Group("/api/v1/products")
	{
		products.GET("/", productHandler.GetAllProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
}
