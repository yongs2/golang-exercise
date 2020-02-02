package main

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/profile", func(c *gin.Context) {
		var form ProfileForm
		if err := c.ShouldBind(&form); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		fmt.Printf("form.Name=[%s], File(%d)[%s]\n", form.Name, form.Avatar.Size, form.Avatar.Filename)
		err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}
		c.JSON(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
