package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		component := hello("Simon")
		component.Render(c.Request.Context(), c.Writer)
	})

	r.Run(":8080")
}
