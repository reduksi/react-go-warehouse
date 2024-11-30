package product

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo ProductRepository
}

func NewHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Master_Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.repo.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}