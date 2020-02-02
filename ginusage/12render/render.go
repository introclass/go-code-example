package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	var msg struct {
		Name    string `json:"user"`
		Message string `json:"message"`
		Number  int    `json: "number"`
	}
	msg.Name = "lijiao，<b>中国</b>"
	msg.Message = "hello"

	route.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, msg)
	})

	route.GET("/securejson", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, msg)
	})

	// curl  "127.0.0.1:8080/jsonp?callback=x"
	//  x({"user":"lijiao","message":"hello","Number":0});%
	route.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, msg)
	})

	route.GET("/asciijson", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, msg)
	})

	route.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, msg)
	})

	route.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hello", "number": 10})
	})

	route.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, msg)
	})

	route.GET("/protobuf", func(c *gin.Context) {
		c.ProtoBuf(http.StatusOK, msg)
	})

	route.Run()
}
