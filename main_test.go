package main

import (
	"bytes"
	"comics-store/config"
	"comics-store/handlers"
	"comics-store/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() {
	dsn := "host=localhost user=postgres password=Darveshova2001 dbname=comics_store port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL")
	}
	db.AutoMigrate(&models.Author{}, &models.Category{}, &models.Comic{}, &models.User{})
	config.DB = db
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/authors", handlers.CreateAuthor)
	r.POST("/categories", handlers.CreateCategory)
	return r
}

func TestCreateAuthor(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	author := models.Author{Name: "Test Author"}
	body, _ := json.Marshal(author)

	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Author")
}

func TestCreateCategory(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	category := models.Category{Name: "Test Category"}
	body, _ := json.Marshal(category)

	req, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Category")
}

func TestInvalidAuthorInput(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	author := models.Author{} 
	body, _ := json.Marshal(author)

	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestInvalidCategoryInput(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	category := models.Category{} 
	body, _ := json.Marshal(category)

	req, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
