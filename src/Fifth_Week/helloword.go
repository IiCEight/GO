package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(r *gin.Context) {
	r.String(http.StatusOK, "Hello world!")
}

func main() {
	r := gin.Default()
	r.GET("/index", index)
	r.Run("127.0.0.1:8080")
}
