package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return route
}

func main() {
	route := setupRoute()
	route.Run()
}
