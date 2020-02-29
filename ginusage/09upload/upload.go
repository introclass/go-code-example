package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// curl -XPOST http://localhost:8080/upload/one -H "Content-Type: multipart/form-data" -F "file=@/tmp/file.txt"
	router.POST("/upload/one", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "request error: %s", err.Error())
			return
		}
		if err := c.SaveUploadedFile(file, "/tmp/"+file.Filename+".saved"); err != nil {
			c.String(http.StatusInternalServerError, "internal error: %s", err.Error())
			return
		}
		c.String(http.StatusOK, "received: %s", file.Filename)
	})

	// curl -XPOST http://localhost:8080/upload/many -H "Content-Type: multipart/form-data" -F "upload[]=@/tmp/file1.txt"  -F "upload[]=@/tmp/file2.txt"
	router.POST("/upload/many", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "request error: %s", err.Error())
			return
		}
		files := form.File["upload[]"]

		received := gin.H{}
		failed := gin.H{}

		for _, file := range files {
			if err := c.SaveUploadedFile(file, "/tmp/"+file.Filename+".saved"); err != nil {
				failed[file.Filename] = err.Error()
				continue
			}
			received[file.Filename] = "ok"
		}
		c.JSON(http.StatusOK, gin.H{
			"received": received,
			"failed":   failed,
		})
	})

	router.Run()
}
