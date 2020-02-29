package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func stringResp(c *gin.Context) {
	c.String(http.StatusOK, "This is a string")
	return
}

type UserInfo struct {
	Name     string            `json:"name"`
	Age      int               `json:"age"`
	Messages []string          `json:"messages"`
	Attr     map[string]string `json:"attr"`
}

var userInfo *UserInfo

func jsonResp(c *gin.Context) {
	c.JSON(http.StatusOK, userInfo)
	return
}

func yamlResp(c *gin.Context) {
	c.YAML(http.StatusOK, userInfo)
	return
}
func xmlResp(c *gin.Context) {
	c.XML(http.StatusOK, userInfo)
	return
}

//func protobufResp(c *gin.Context) {
//	c.ProtoBuf(http.StatusOK, userInfo)
//	return
//}

func main() {
	userInfo = &UserInfo{
		Name:     "xiaoming",
		Age:      13,
		Messages: []string{"Welcome1", "Welcome2", "Welcome3"},
		Attr: map[string]string{
			"hobby":   "footbool",
			"country": "china",
		},
	}

	r := gin.Default()

	r.GET("/string", stringResp)
	r.GET("/json", jsonResp)
	r.GET("/yaml", yamlResp)
	r.GET("/xml", xmlResp)
	//TODO:
	// protobuf
	// header
	// data
	// file
	// redirect
	r.Run()
}
