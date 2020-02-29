// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat: lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required" `
}

func bindJSON(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, login)
}

func bindXML(c *gin.Context) {
	var login Login
	if err := c.ShouldBindXML(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, login)
}

func bindALL(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, login)
}

func main() {
	router := gin.Default()

	// 几个 modelbind 函数的用法
	// c.ShouldBindQuery()  只解析 GET 参数
	// c.ShouldBind()       解析 GET 和 POST 数据
	// c.ShouldBindUri()    解析路径参数
	// c.ShouldBindHeader() 解析请求头

	// curl -X POST 127.0.0.1:8080/loginjson -d '{"user": "lijiao", "password":"123"}'
	router.POST("/loginjson", bindJSON)

	// curl -X POST 127.0.0.1:8080/loginxml \
	//    -d '<?xml version="1.0" encoding="UTF-8"?><root><user>user</user><password>123</password> </root>'
	router.POST("/loginxml", bindXML)

	// c.ShouldBind() 根据 content-type 选择合适的反序列化方式
	// curl -X POST 127.0.0.1:8080/loginform -H 'content-type: application/xml' \
	//      -d '<?xml version="1.0" encoding="UTF-8"?><root><user>user</user><password>123</password> </root>'
	// curl -X POST 127.0.0.1:8080/loginform -H 'content-type: application/json' \
	//      -d '{"user": "lijiao", "password":"123"}'
	// curl -X POST 127.0.0.1:8080/loginform -d "user=lijiao&password=123"
	router.POST("/loginform", bindALL)

	router.Run(":8080")
}
