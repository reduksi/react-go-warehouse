package routes

import (
	"warehouse/internal/product"

	"github.com/gin-gonic/gin"
)

func productRoutes(router *gin.Engine, handler *product.ProductHandler) {
	router.GET("/products", handler.GetAllProducts)
	router.POST("/products", handler.CreateProduct)
}