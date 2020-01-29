package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
