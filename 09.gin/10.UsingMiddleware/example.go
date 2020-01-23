package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
	router.Run(":8080")
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Benmark Middleware 1.\n")
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Auth Middleware 1.\n")
	}
}

func benchEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Benmark Middleware 2.")
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

func analyticsEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Auth Middleware Analytics Point.")
}
