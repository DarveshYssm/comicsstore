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


	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)


	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	
		protected.GET("/comics", handlers.GetAllComics)
		protected.POST("/comics", handlers.CreateComic)
		protected.GET("/comics/:id", handlers.GetComicByID)
		protected.PUT("/comics/:id", handlers.UpdateComic)
		protected.DELETE("/comics/:id", handlers.DeleteComic)

	
		protected.POST("/authors", handlers.CreateAuthor)

		
		protected.POST("/categories", handlers.CreateCategory)
		protected.GET("/categories", handlers.GetAllCategories)


	router.Run(":8080")
}
