package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		defer file.Close()
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func main() {
	route := gin.Default()

	t, err := loadTemplate()
	if err != nil {
		log.Fatal("load tempalte: ", err.Error())
	}

	route.SetHTMLTemplate(t)

	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", nil)
	})
	route.Run()
}
