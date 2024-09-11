package routes

import (
	"library-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthorRoutes(r *gin.Engine, db *gorm.DB) {
	authorRoutes := r.Group("/authors")
	{
		authorRoutes.POST("/", createAuthor(db))
		authorRoutes.GET("/", getAuthors(db))
		authorRoutes.GET("/:id", getAuthor(db))
		authorRoutes.PUT("/:id", updateAuthor(db))
		authorRoutes.DELETE("/:id", deleteAuthor(db))
	}
}

func createAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author models.Author
		if err := c.ShouldBindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&author)
		c.JSON(http.StatusCreated, author)
	}
}

func getAuthors(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authors []models.Author
		db.Find(&authors)
		c.JSON(http.StatusOK, authors)
	}
}

func getAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var author models.Author
		if err := db.First(&author, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		c.JSON(http.StatusOK, author)
	}
}

func updateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var author models.Author
		if err := db.First(&author, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		if err := c.ShouldBindJSON(&author); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&author)
		c.JSON(http.StatusOK, author)
	}
}

func deleteAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Author{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
