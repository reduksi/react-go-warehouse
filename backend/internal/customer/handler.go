package customer

import (
	"net/http"
	"warehouse/models"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	repo CustomerRepository
}

func NewHandler(repo CustomerRepository) *CustomerHandler {
	return &CustomerHandler{repo: repo}
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.repo.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer models.Master_Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCustomer, err := h.repo.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCustomer)
}