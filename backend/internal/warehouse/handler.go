package warehouse

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	repo WarehouseRepository
}

func NewHandler(repo WarehouseRepository) *WarehouseHandler {
	return &WarehouseHandler{repo: repo}
}

func (h *WarehouseHandler) GetAllWarehouses(c *gin.Context) {
	warehouses, err := h.repo.GetAllWarehouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch warehouses"})
		return
	}

	c.JSON(http.StatusOK, warehouses)
}

func (h *WarehouseHandler) CreateWarehouse(c *gin.Context) {
	var warehouse models.Master_Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdWarehouse, err := h.repo.CreateWarehouse(warehouse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdWarehouse)
}