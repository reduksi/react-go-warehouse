package main

import (
	"warehouse/cmd/routes"
	"warehouse/db"
	"warehouse/internal/customer"
	"warehouse/internal/product"
	"warehouse/internal/supplier"
	"warehouse/internal/transaction/issueTransaction"
	"warehouse/internal/transaction/receiptTransaction"
	"warehouse/internal/warehouse"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://your-frontend-url.com"}, // Add your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dbConn := db.InitDB()

	supplierRepo := supplier.NewRepository(dbConn)
	productRepo := product.NewRepository(dbConn)
	warehouseRepo := warehouse.NewRepository(dbConn)
	receiptTransactionRepo := receiptTransaction.NewRepository(dbConn)
	issueTransactionRepo := issueTransaction.NewRepository(dbConn)
	customerRepo := customer.NewRepository(dbConn)

	supplierHandler := supplier.NewHandler(supplierRepo)
	productHandler := product.NewHandler(productRepo)
	warehouseHandler := warehouse.NewHandler(warehouseRepo)
	receiptTransactionHandler := receiptTransaction.NewHandler(receiptTransactionRepo)
	issueTransactionHandler := issueTransaction.NewHandler(issueTransactionRepo)
	customerHandler := customer.NewHandler(customerRepo)

	routes.SetupRoutes(router, supplierHandler, productHandler, warehouseHandler, receiptTransactionHandler, issueTransactionHandler, customerHandler)

	router.Run(":8080")
}