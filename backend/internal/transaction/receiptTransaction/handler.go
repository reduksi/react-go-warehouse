package receiptTransaction

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	repo TransactionRepository
}

func NewHandler(repo TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) GetAllGoodsReceiptHeaders(c *gin.Context) {
	headers, err := h.repo.GetAllGoodsReceiptHeaders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch goods receipt headers"})
		return
	}

	c.JSON(http.StatusOK, headers)
}
func (h *TransactionHandler) CreateGoodsReceipt(c *gin.Context) {
	var request struct {
		Header  models.GoodsReceiptHeader        `json:"header"`
		Details []models.GoodsReceiptDetail     `json:"details"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdHeader, err := h.repo.CreateGoodsReceipt(request.Header, request.Details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdHeader)
}