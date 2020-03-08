package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		tNow := time.Now()
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! in path " + cCp.Request.URL.Path)
			log.Println("Time=", time.Since(tNow))
		}()
		c.JSON(http.StatusOK, gin.H{"status": c.Request.URL.Path})
	})

	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! in path " + c.Request.URL.Path)
		c.JSON(http.StatusOK, gin.H{"status": c.Request.URL.Path})
	})

	router.Run(":8080")
}
