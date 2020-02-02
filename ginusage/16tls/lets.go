package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	autotls.Run(r, "example1.com", "example2.com")
}
