package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	//指定校验函数： bookabledate
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname""`
	Email     string `bingding: "required,email"`
}

func bookabledate(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	// 有问题，没通过，待查
	route := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//master分支中的 func 定义发生了变换，这里是 1.5.0 中的用法
		v.RegisterValidation("bookabledate", bookabledate)
		//Struct 级别的校验，该写法对于 1.5.0 不生效
		v.RegisterStructValidation(func(v *validator.Validate, structLevel *validator.StructLevel) {
			user := structLevel.CurrentStruct.Interface().(User)

			structLevel.ReportError(reflect.ValueOf(user.FirstName), "FirstName", "fname", "fnameorlname")
			if len(user.FirstName) == 0 && len(user.LastName) == 0 {
				structLevel.ReportError(reflect.ValueOf(user.FirstName), "FirstName", "fname", "fnameorlname")
				structLevel.ReportError(reflect.ValueOf(user.LastName), "LastName", "lname", "fnameorlname")
			}
		}, &User{})
	}

	route.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User validation failed",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, &u)
	})
	route.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBindWith(&b, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
	route.Run()
}
