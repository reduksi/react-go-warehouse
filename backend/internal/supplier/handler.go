package supplier

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
    return &Handler{Repo: repo}
}

func (h *Handler) GetAllSuppliers(c *gin.Context) {
    suppliers, err := h.Repo.GetAllSuppliers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, suppliers)
}

func (h *Handler) CreateSupplier(c *gin.Context) {
	var supplier models.Master_Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSupplier, err := h.Repo.CreateSupplier(supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSupplier)
}