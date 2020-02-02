package main

import "github.com/gin-gonic/gin"

func echo(c *gin.Context) {
	c.JSON(200, gin.H{
		"method": c.Request.Method,
		"uri":    c.Request.URL.String(),
	})
}

func main() {
	router := gin.Default()

	router.GET("/get", echo)
	router.POST("/post", echo)
	router.PUT("/put", echo)
	router.DELETE("/delete", echo)
	router.PATCH("/patch", echo)
	router.HEAD("/head", echo)
	router.OPTIONS("/option", echo)
	router.Run()
}
