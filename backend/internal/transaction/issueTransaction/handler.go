package issueTransaction

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

func (h *TransactionHandler) GetAllGoodsIssueHeaders(c *gin.Context) {
	headers, err := h.repo.GetAllGoodsIssueHeaders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch goods issue headers"})
		return
	}

	c.JSON(http.StatusOK, headers)
}

func (h *TransactionHandler) CreateGoodsIssue(c *gin.Context) {
	var request struct {
		Header  models.GoodsIssueHeader        `json:"header"`
		Details []models.GoodsIssueDetail     `json:"details"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdHeader, err := h.repo.CreateGoodsIssueHeader(request.Header, request.Details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdHeader)
}