package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s\n", name)
}

func userNameAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	c.String(http.StatusOK, "%s is %s\n", name, action)
}

type UserInfo struct {
	Name   string `uri:"name"`
	Action string `uri:"action"`
}

func bindURI(c *gin.Context) {
	var userInfo UserInfo
	c.BindUri(&userInfo)
	c.JSON(http.StatusOK, userInfo)
}

func main() {
	router := gin.Default()

	// curl  127.0.0.1:8080/user/xiaoming
	// :name 必须存在
	// 匹配 /user/NAME
	// 不匹配 /user 或者 /user/
	router.GET("/user/:name", userName)

	// curl 127.0.0.1:8080/user/xiaoming/run
	// *action 是可选的
	// 匹配 /user/NAME/
	// 匹配 /user/NAME/ACTION
	// 如果 /user/NAME 没有对应的 router，重定向到 /user/NAME/
	router.GET("/user/:name/*action", userNameAction)

	// curl -X POST 127.0.0.1:8080/user/xiaoming/run
	// POST 也可以使用路径参数
	router.POST("/user/:name/*action", userNameAction)

	// curl 127.0.0.1:8080/bind/user/xiaoming/run
	router.GET("/bind/user/:name/*action", bindURI)

	router.Run()
}
