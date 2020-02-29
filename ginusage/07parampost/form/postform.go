package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postForm(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "guest")
	c.String(http.StatusOK, "%s %s", name, message)
}

func postFormMap(c *gin.Context) {
	ids := c.PostFormMap("ids")
	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func postFormArray(c *gin.Context) {
	ids := c.PostFormArray("ids")
	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

type UserInfo struct {
	Name     string            `form:"name"`
	Age      int               `form:"age"`
	Message  []string          `form:"message"`
	MoreInfo map[string]string `form:"moreinfo"`
}

func postBind(c *gin.Context) {
	var userInfo UserInfo
	c.ShouldBind(&userInfo)
	//gin 1.5 的 ShouldBind 似乎不支持 map，单独读取 map
	userInfo.MoreInfo = c.PostFormMap("moreinfo")
	c.JSON(http.StatusOK, userInfo)
}

func main() {
	router := gin.Default()

	// curl -XPOST 127.0.0.1:8080/form -d "name=xiaoming&message=welcome!"
	router.POST("/form", postForm)

	// curl -XPOST 127.0.0.1:8080/map -d "ids[first]=100&ids[second]=200"
	router.POST("/map", postFormMap)

	// curl -XPOST 127.0.0.1:8080/map -d "ids=100&ids=200"
	router.POST("/arr", postFormArray)

	// curl -XPOST 127.0.0.1:8080/bind  -d "name=xiaoming&age=11&message=hello1&message=hello2&moreinfo[country]=china&moreinfo[hobby]=football"
	router.POST("/bind", postBind)

	router.Run()
}
