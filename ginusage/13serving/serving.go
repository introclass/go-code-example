package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	route.Static("/static", "/tmp/static")
	route.StaticFS("/static2", http.Dir("/tmp/static2"))
	route.StaticFile("/file", "/tmp/static/file.txt")

	route.GET("/stream", func(c *gin.Context) {
		resp, err := http.Get("https://www.lijiaocn.com/money/img/self.jpg")
		if err != nil || resp.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := resp.Body
		contentLength := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	route.Run()
}
