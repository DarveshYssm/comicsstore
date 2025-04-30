package main

import (
	"comics-store/config"
	"comics-store/handlers"
	"comics-store/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()


	router := gin.Default()


	router.POST("/api/register", handlers.Register)
	router.POST("/api/login", handlers.Login)


	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/comics", handlers.GetAllComics)
	protected.POST("/comics", handlers.CreateComic)
	protected.GET("/comics/:id", handlers.GetComicByID)
	protected.PUT("/comics/:id", handlers.UpdateComic)
	protected.DELETE("/comics/:id", handlers.DeleteComic)


	protected.POST("/authors", handlers.CreateAuthor)
	protected.GET("/authors", handlers.GetAllAuthors)
	protected.GET("/authors/:id", handlers.GetAuthorByID)
	protected.PUT("/authors/:id", handlers.UpdateAuthor)
	protected.DELETE("/authors/:id", handlers.DeleteAuthor)


	protected.POST("/categories", handlers.CreateCategory)
	protected.GET("/categories", handlers.GetAllCategories)
	protected.GET("/categories/:id", handlers.GetCategoryByID)
	protected.PUT("/categories/:id", handlers.UpdateCategory)
	protected.DELETE("/categories/:id", handlers.DeleteCategory)

	router.Run(":8080")
}
