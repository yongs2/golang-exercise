package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	})

	log.Fatal(autotls.Run(router, "127-0-0-1.sslip.io", "172-17-0-2.sslip.io"))
}
