package routes

import (
	"library-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCustomerRoutes(r *gin.Engine, db *gorm.DB) {
	customerRoutes := r.Group("/customers")
	{
		customerRoutes.POST("/", createCustomer(db))
		customerRoutes.GET("/", getCustomers(db))
		customerRoutes.GET("/:id", getCustomer(db))
		customerRoutes.PUT("/:id", updateCustomer(db))
		customerRoutes.DELETE("/:id", deleteCustomer(db))
	}
}

func createCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer models.Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&customer)
		c.JSON(http.StatusCreated, customer)
	}
}

func getCustomers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customers []models.Customer
		db.Find(&customers)
		c.JSON(http.StatusOK, customers)
	}
}

func getCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var customer models.Customer
		if err := db.First(&customer, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}

func updateCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var customer models.Customer
		if err := db.First(&customer, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&customer)
		c.JSON(http.StatusOK, customer)
	}
}

func deleteCustomer(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Customer{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
