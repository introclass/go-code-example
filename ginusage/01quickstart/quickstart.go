// Create: 2019/07/19 17:46:00 Change: 2019/07/19 17:47:00
// FileName: modelbind.go
// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com> wechat:lijiaocn
//
// Distributed under terms of the GPL license.

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//r.Run() // listen and serve on 0.0.0.0:8080
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
