package main

import (
	_ "Fifth_Week/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Saber
// @version 1.0
// @description welcome!
// @termsOfService http://arturia.top

// @contact.name iSaber
// @contact.url http://arturia.top
// @contact.email 1346959878@qq.com

// @host localhost
// @BasePath D:\Project\Visual Studio Code\Go
func main() {
	r := gin.Default()
	r.GET("/index", gs.WrapHandler(swaggerFiles.Handler))
	// r.GET("/index", index)
	r.Run()
}

// index 升级版帖子列表接口
// @Summary 升级版帖子列表接口

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "<h1> Hello, world!</h1>", nil)

}
