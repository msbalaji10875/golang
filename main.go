package main

import (
	"os"

	"example.com/golang-jwt-project/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(context *gin.Context) {
		context.JSON(200, gin.H{"success": "Access granted for API 1"})
	})

	router.Run(":" + port)
}
