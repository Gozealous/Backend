package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "This is the startup page!")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.GET("/apple", func(c *gin.Context) {
		c.String(200, "I'm a small apple!")
	})

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
