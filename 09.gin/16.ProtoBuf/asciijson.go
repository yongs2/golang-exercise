package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(c *gin.Context) {
		data := gin.H{
			"lang": "go헬로우",
			"tag" : "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
