package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" 	 json:"user" 	 xml:"user"  	binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func GetBody(c *gin.Context) string {
	buf := make([]byte, 1024)
	num, _ := c.Copy().Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	return reqBody
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		reqBody := GetBody(c)
		fmt.Printf("Body=[%s]\n", reqBody)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))	// Write body back

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("user:[%s], passwd:[%s]\n", json.User, json.Password)
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login

		reqBody := GetBody(c)
		fmt.Printf("Body=[%s]\n", reqBody)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))	// Write body back

		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("user:[%s], passwd:[%s]\n", xml.User, xml.Password)

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login

		reqBody := GetBody(c)
		fmt.Printf("Body=[%s]\n", reqBody)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))	// Write body back

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("user:[%s], passwd:[%s]\n", form.User, form.Password)

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}
