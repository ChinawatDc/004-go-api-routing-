package main

import (
	"log"
	"net/http"

	"go-api-routing/middleware"
	"go-api-routing/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register middleware
	r.Use(middleware.LoggerMiddleware())

	// Default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to go-api-routing!"})
	})

	// Register user routes
	routes.RegisterUserRoutes(r)

	log.Println("🚀 Server is running at http://localhost:8080")
	r.Run(":8080")
}
