package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//匹配 /user/NAME，不能匹配 /user 或者 /user/
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//匹配 /user/NAME/
	//匹配 /user/NAME/ACTION
	//如果 /user/NAME 没有对应的 router，重定向到 /user/NAME/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "%s %s", name, action)
	})

	// gin 1.5.0 开始，c.FullPath()保存完整的 uri 定义
	router.POST("/user/:name/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "fullpath is: %s", c.FullPath())
	})

	router.Run()
}
