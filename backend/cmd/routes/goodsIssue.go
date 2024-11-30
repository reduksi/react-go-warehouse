package routes

import (
	"warehouse/internal/transaction/issueTransaction"

	"github.com/gin-gonic/gin"
)

func goodsIssueRoutes(router *gin.Engine, handler *issueTransaction.TransactionHandler) {
	router.GET("/goods-issue", handler.GetAllGoodsIssueHeaders)
	router.POST("/goods-issue", handler.CreateGoodsIssue)
}