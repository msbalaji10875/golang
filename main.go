package main

import (
	"os"

	"example.com/golang-jwt-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes()
	routes.UserRoutes()

	router.GET("/api-1", func(context *gin.Context) {
		context.JSON(200, gin.H{"success": "Access grabted for API 1"})
	})

	router.Run(":" + port)
}
