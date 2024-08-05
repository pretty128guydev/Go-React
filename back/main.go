package main

import (
	"back/db"
	"back/models"
	"back/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Auto migrate your models
	db.DB.AutoMigrate(&models.User{})

	router := gin.Default()

	// Allow CORS requests from frontend
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Add a root endpoint for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	routes.AuthRoutes(router)
	router.Run(":8080")
}
