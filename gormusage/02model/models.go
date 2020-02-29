package main

import (
	"go-code-example/gormusage/common"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "gorm:password123@(localhost)/gormdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal(err.Error())
	}
	defer db.Close()

	// 表名默认是 struct 名称的复数形式，使用单数形式：
	db.SingularTable(true)

	// 自定义表名命名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	if !db.HasTable(&common.User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&common.User{})
	}

	if !db.HasTable(&common.TableStudent{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&common.TableStudent{})
	}
}
