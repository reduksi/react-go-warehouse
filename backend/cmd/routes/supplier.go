package routes

import (
	"warehouse/internal/supplier"

	"github.com/gin-gonic/gin"
)

func supplierRoutes(router *gin.Engine, handler *supplier.Handler) {
	router.GET("/suppliers", handler.GetAllSuppliers)
	router.POST("/suppliers", handler.CreateSupplier)
}