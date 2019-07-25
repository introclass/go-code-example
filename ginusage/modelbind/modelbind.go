// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type A struct {
	Value int `json:"value" form:"value" xml:"value"`
}

var a A

func SetA(c *gin.Context) {
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, a)
}

func GetA(c *gin.Context) {
	c.JSON(200, a)
}

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/a", GetA)
		v1.POST("/a", SetA)
	}

	router.Run(":8080")
}
