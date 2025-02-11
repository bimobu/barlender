package main

import (
	"barlender/components/layout"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		component := layout.Template("Simon")
		component.Render(c.Request.Context(), c.Writer)
	})

	r.Run(":8080")
}
