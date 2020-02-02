package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.LoadHTMLFiles("/tmp/templates/index.tmpl")
	route.LoadHTMLGlob("/tmp/templates/**/*")
	route.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "lijiaocn.com",
		})
	})
	route.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "lijiaocn.com",
		})
	})
	route.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://wwww.lijiaocn.com")
	})
	route.Run()
}
