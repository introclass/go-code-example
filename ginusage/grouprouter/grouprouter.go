// Create: 2019/07/22 10:50:00 Change: 2019/07/22 11:05:15
// FileName: waitgroup.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import "github.com/gin-gonic/gin"

type A struct {
	Value int
}

var a A

func SetA(c *gin.Context) {
	a.Value = 3
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
