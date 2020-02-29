package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CustormMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next() //继续执行其它中间件
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Println(latency, status)
	}
}

func main() {
	router := gin.New()

	//全局中间件：日志写入 gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	router.Use(CustormMiddleWare())

	//全局中间件：如果发生 panic 恢复，并回复 500
	router.Use(gin.Recovery())

	//作用于当前路由的中间件
	router.GET("/echo", func(c *gin.Context) {
		c.Set("middleware1", "ok")
	}, func(c *gin.Context) {
		c.Set("middleware2", "ok")
	}, func(c *gin.Context) {
		resp := gin.H{}
		if v, ok := c.Get("middleware1"); ok {
			resp["middleware1"] = v
		}
		if v, ok := c.Get("middleware2"); ok {
			resp["middleware2"] = v
		}
		resp["path"] = "/echo"
		c.JSON(http.StatusOK, resp)
	})

	//作用于 group 的中间件
	root := router.Group("/")
	root.Use(func(c *gin.Context) {
		c.Set("middleware3", "ok")
	})
	{
		root.GET("/a", func(c *gin.Context) {
			resp := gin.H{}
			if v, ok := c.Get("middleware3"); ok {
				resp["middleware3"] = v
			}
			resp["path"] = "/a"
			c.JSON(http.StatusOK, resp)
		})

		root.GET("/b", func(c *gin.Context) {
			resp := gin.H{}
			if v, ok := c.Get("middleware3"); ok {
				resp["middleware3"] = v
			}
			resp["path"] = "/b"
			c.JSON(http.StatusOK, resp)
		})
	}

	router.Run()
}
