package main

import (
	"gin-gonic/pkg/hello"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, hello.GetMessage())
	})

	r.Run(":8080")
}
