package routes

import (
	"warehouse/internal/transaction/receiptTransaction"

	"github.com/gin-gonic/gin"
)

func goodsReceiptRoutes(router *gin.Engine, handler *receiptTransaction.TransactionHandler) {
	router.GET("/stock-report", handler.GetStockReport)
	router.GET("/goods-receipt", handler.GetAllGoodsReceiptHeaders)
	router.POST("/goods-receipt", handler.CreateGoodsReceipt)
}