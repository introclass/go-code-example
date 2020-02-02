package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//匹配：127.0.0.1:8080/query?name=XX&age=11
	router.GET("/query", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		age := c.Query("age")

		c.String(http.StatusOK, "%s %s", name, age)
	})

	//匹配："127.0.0.1:8080/map?ids[a]=1234&ids[b]=hello"
	router.GET("/map", func(c *gin.Context) {
		ids := c.QueryMap("ids")

		c.JSON(http.StatusOK, gin.H{
			"ids": ids,
		})
	})

	router.Run()
}
