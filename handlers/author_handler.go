package handlers

import (
	"comics-store/config"
	"comics-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
func AuthorHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Authors endpoint is working"))
}

func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	result := config.DB.First(&author, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var newAuthor models.Author

	if err := c.ShouldBindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newAuthor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author name is required"})
		return
	}

	result := config.DB.Create(&newAuthor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add author"})
		return
	}

	c.JSON(http.StatusCreated, newAuthor)
}

func GetAllAuthors(c *gin.Context) {
	var authors []models.Author

	result := config.DB.Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := config.DB.Save(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update author"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	result := config.DB.Delete(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
