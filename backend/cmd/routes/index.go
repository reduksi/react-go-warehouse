package routes

import (
	"warehouse/internal/customer"
	"warehouse/internal/product"
	"warehouse/internal/supplier"
	"warehouse/internal/transaction/issueTransaction"
	"warehouse/internal/transaction/receiptTransaction"
	"warehouse/internal/warehouse"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, supplierHandler *supplier.Handler, productHandler *product.ProductHandler, warehouseHandler *warehouse.WarehouseHandler,
	receiptTransactionHandler *receiptTransaction.TransactionHandler, issueTransactionHandler *issueTransaction.TransactionHandler, customerHandler *customer.CustomerHandler) {

	supplierRoutes(router, supplierHandler)
	productRoutes(router, productHandler)
	warehouseRoutes(router, warehouseHandler)
	goodsReceiptRoutes(router, receiptTransactionHandler)
	goodsIssueRoutes(router, issueTransactionHandler)
	customerRoutes(router, customerHandler)
}