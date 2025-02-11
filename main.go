package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	router := gin.Default()

	// Define a route for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
