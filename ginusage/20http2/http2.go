package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>Https Test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
<h1 style="color:red;">Welcome {{.user}} </h1>
</body>
</html>
`))

func main() {
	route := gin.Default()
	route.Static("/assets", "./assets")
	route.SetHTMLTemplate(html)

	route.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %s", err.Error())
			}
		}
		c.HTML(http.StatusOK, "https", gin.H{"user": "guest"})
	})
	route.Run()
}
