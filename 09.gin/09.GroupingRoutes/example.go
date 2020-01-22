package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login"})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "submit"})
}

func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{"message": "read"})
}
