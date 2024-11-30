package routes

import (
	"warehouse/internal/warehouse"

	"github.com/gin-gonic/gin"
)

func warehouseRoutes(router *gin.Engine, handler *warehouse.WarehouseHandler) {
	router.GET("/warehouses", handler.GetAllWarehouses)
	router.POST("/warehouses", handler.CreateWarehouse)
}