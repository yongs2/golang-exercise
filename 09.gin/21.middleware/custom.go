package main

import (
	"log"
	//"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()

		latency := time.Since(t)
		log.Println("latency=", latency)

		status := c.Writer.Status()
		log.Println("status=", status)
	}
}

func main() {
	router := gin.Default()
	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println("example=", example)
		//c.JSON(http.StatusOK, gin.H{"example": example})
	})

	router.Run(":8080")
}
