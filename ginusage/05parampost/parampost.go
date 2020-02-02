package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// curl -XPOST 127.0.0.1:8080/form -d "name=manu&message=this_is_great"
	router.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "guest")
		c.String(http.StatusOK, "%s %s", name, message)
	})

	// curl -XPOST 127.0.0.1:8080/map -d "names[first]=thinkerou&names[second]=tianou"
	router.POST("/map", func(c *gin.Context) {
		names := c.PostFormMap("names")
		c.JSON(http.StatusOK, gin.H{
			"names": names,
		})
	})

	router.Run()
}
