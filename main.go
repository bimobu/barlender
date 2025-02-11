// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Create a default gin router
// 	router := gin.Default()

// 	// Define a route for the root path
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello, World!",
// 		})
// 	})

// 	// Run the server on port 8080
// 	router.Run(":8080")
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("askfjhbajbseljbhfa")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
