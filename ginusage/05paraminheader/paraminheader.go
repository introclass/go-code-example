package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func header(c *gin.Context) {
	token := c.GetHeader("TOKEN")
	c.JSON(http.StatusOK, gin.H{
		"TOKEN": token,
	})
}

type Token struct {
	Token string `header:"token"`
}

func headerBind(c *gin.Context) {
	var token Token
	c.BindHeader(&token)
	c.JSON(http.StatusOK, token)
}

func main() {
	r := gin.Default()
	// curl -H "TOKEN: 123456" 127.0.0.1:8080/header
	r.GET("/header", header)
	// curl -H "TOKEN: 123456" 127.0.0.1:8080/header/bind
	r.GET("/header/bind", headerBind)
	r.Run()
}
