package main

import "github.com/gin-gonic/gin"

type application struct {
	Router *gin.Engine
}

func newApp() *application {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return &application{Router: router}
}

func (a *application) Start() {
	a.Router.Run()
}

func main() {
	newApp().Start()
}
