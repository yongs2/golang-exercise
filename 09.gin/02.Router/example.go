package main

import "github.com/gin-gonic/gin"

func getting(c *gin.Context)  { c.JSON(200, gin.H{"message": "getting"}) }
func posting(c *gin.Context)  { c.JSON(200, gin.H{"message": "posting"}) }
func putting(c *gin.Context)  { c.JSON(200, gin.H{"message": "putting"}) }
func deleting(c *gin.Context) { c.JSON(200, gin.H{"message": "deleting"}) }
func patching(c *gin.Context) { c.JSON(200, gin.H{"message": "patching"}) }
func head(c *gin.Context)     { c.JSON(200, gin.H{"message": "head"}) }
func options(c *gin.Context)  { c.JSON(200, gin.H{"message": "options"}) }

func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	router.Run()
}
