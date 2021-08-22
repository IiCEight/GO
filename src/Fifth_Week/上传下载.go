package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "load.html", gin.H{})
}

func up(c *gin.Context) {
	upfile, err := c.FormFile("myfile")
	checkerr(err)
	c.SaveUploadedFile(upfile, upfile.Filename)
}

func down(c *gin.Context) {
	path := "D:\\Project\\Visual Studio Code\\Go\\src\\Fifth_Week\\猫耳呆毛two.png"
	c.File(path)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("temp/*")
	r.GET("/index", index)
	r.POST("/up", up)
	r.GET("/down", down)
	r.Run()
}
