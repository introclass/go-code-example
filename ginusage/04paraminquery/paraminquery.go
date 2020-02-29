package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func query(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	age := c.Query("age")

	c.String(http.StatusOK, "%s is %s years\n", name, age)
}

func queryArray(c *gin.Context) {
	ids := c.QueryArray("ids")
	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func queryMap(c *gin.Context) {
	ids := c.QueryMap("ids")
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

func queryBind(c *gin.Context) {
	var userInfo UserInfo
	c.BindQuery(&userInfo)
	//gin 1.5 的 BindQuery 似乎不支持 map，单独获取
	userInfo.MoreInfo = c.QueryMap("moreinfo")
	c.JSON(http.StatusOK, userInfo)
}

func main() {
	router := gin.Default()
	// curl 127.0.0.1:8080/query?name=xiaoming&age=11
	router.GET("/query", query)
	// curl "127.0.0.1:8080/arr?ids=A&ids=B"
	router.GET("/arr", queryArray)
	// curl "127.0.0.1:8080/map?ids[a]=1234&ids[b]=hello"
	router.GET("/map", queryMap)
	// curl "127.0.0.1:8080/bind?name=xiaoming&age=11&message=hello1&message=hello2&moreinfo[country]=china&moreinfo[hobby]=football"
	router.GET("/bind", queryBind)

	router.Run()
}
