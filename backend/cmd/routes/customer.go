package routes

import (
	"warehouse/internal/customer"

	"github.com/gin-gonic/gin"
)

func customerRoutes(router *gin.Engine, handler *customer.CustomerHandler) {
	router.GET("/customers", handler.GetAllCustomers)
	router.POST("/customers", handler.CreateCustomer)
}