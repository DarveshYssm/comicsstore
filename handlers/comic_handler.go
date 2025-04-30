package handlers

import (
	"comics-store/config"
	"comics-store/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComicHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Comics endpoint is working"))
}


func GetAllComics(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	authorID := c.DefaultQuery("author_id", "")
	categoryID := c.DefaultQuery("category_id", "")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	var comics []models.Comic
	var total int64

	query := config.DB.Model(&models.Comic{})

	if authorID != "" {
		query = query.Where("author_id = ?", authorID)
	}

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	query.Count(&total)

	result := query.Select("id, title, description, price, author_id, category_id").
		Limit(limitInt).
		Offset(offset).
		Find(&comics)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comics"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       comics,
		"total":      total,
		"page":       pageInt,
		"limit":      limitInt,
		"totalPages": int(math.Ceil(float64(total) / float64(limitInt))),
	})
}

func CreateComic(c *gin.Context) {
	var newComic models.Comic

	if err := c.ShouldBindJSON(&newComic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newComic.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comic title is required"})
		return
	}
	if newComic.Description == "" || newComic.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description and positive price are required"})
		return
	}

	result := config.DB.Create(&newComic)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comic"})
		return
	}

	config.DB.Preload("Author").Preload("Category").First(&newComic, newComic.ID)

	c.JSON(http.StatusCreated, newComic)
}

func GetComicByID(c *gin.Context) {
	id := c.Param("id")
	var comic models.Comic

	result := config.DB.Preload("Author").Preload("Category").First(&comic, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comic not found"})
		return
	}

	c.JSON(http.StatusOK, comic)
}

func UpdateComic(c *gin.Context) {
	id := c.Param("id")
	var comic models.Comic

	if err := config.DB.First(&comic, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comic not found"})
		return
	}

	if err := c.ShouldBindJSON(&comic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if comic.Title == "" || comic.Description == "" || comic.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title, description, and positive price are required"})
		return
	}

	result := config.DB.Save(&comic)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comic"})
		return
	}

	config.DB.Preload("Author").Preload("Category").First(&comic, id)


	c.JSON(http.StatusOK, comic)
}

func DeleteComic(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.Delete(&models.Comic{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comic not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comic deleted successfully"})
}
