package main

import (
	"net/http"

	_ "src/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title rtp cloud
// @description rtp 的云端系统
// @host localhost:8080

func main() {
	r := gin.Default()

	r.GET("/index", index)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

//@Tags index
// @Summary index
// @Description none
// @Produce  json
// @Success 200 {string} string "hsdf"
// @Failure 500 {string} string "no"
// @Router /index [get]
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "hahhah", nil)

}
