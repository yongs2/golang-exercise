package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	router.Run(":8080")
}
