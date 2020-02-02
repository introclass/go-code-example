package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	route := gin.Default()

	route.GET("/async", func(c *gin.Context) {
		cCp := c.Copy() //启动协程，必须重新拷贝一份
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done: " + cCp.Request.URL.Path)
		}()
	})

	route.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done: " + c.Request.URL.Path)
	})

	route.Run()
}
