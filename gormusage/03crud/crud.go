package main

import (
	"fmt"
	"go-code-example/gormusage/common"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func create(db *gorm.DB) error {
	var num string = "00001"
	user := common.User{
		Name:         "user1",
		MemberNumber: &num,
		Role:         "normal",
	}
	return db.Save(&user).Error
}

func create2(db *gorm.DB) error {
	student := &common.Student{
		Name: "xiao a",
	}
	return db.Save(student).Error
}

func read(db *gorm.DB) (*common.User, error) {
	var user common.User
	var num string = "00001"
	dberr := db.Where(&common.User{MemberNumber: &num}).First(&user).Error
	if dberr != nil {
		return nil, dberr
	}
	return &user, nil
}

func read2(db *gorm.DB) (*common.Student, error) {
	var student = &common.Student{
		Name: "xiao a",
	}
	dberr := db.Where(student).First(student).Error
	if dberr != nil {
		return nil, dberr
	}
	return student, nil
}

func update(db *gorm.DB) error {
	var num string = "00002"
	user := common.User{
		Name:         "user1",
		MemberNumber: &num,
		Role:         "admin",
	}
	var oriUser common.User
	dberr := db.Where(&common.User{MemberNumber: user.MemberNumber}).First(&oriUser).Error
	if dberr != nil {
		return dberr
	}
	user.ID = 100
	return db.Save(&user).Error
}

func delete(db *gorm.DB) error {
	var num string = "00001"
	user := common.User{
		MemberNumber: &num,
	}
	return db.Delete(&user).Error
}

func main() {
	db, err := gorm.Open("mysql", "gorm:password123@(localhost)/gormdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal(err.Error())
	}
	defer db.Close()

	//Create
	if err := create(db); err != nil {
		glog.Error(err.Error())
	}

	if err := create2(db); err != nil {
		glog.Error(err.Error())
	}

	//Read
	if user, err := read(db); err != nil {
		glog.Error(err.Error())
	} else {
		fmt.Printf("read user info: %v\n", user)
	}

	if student, err := read2(db); err != nil {
		glog.Error(err.Error())
	} else {
		fmt.Printf("read user info: %v\n", student)
	}

	//Update
	update(db)
	if user, err := read(db); err != nil {
		glog.Error(err.Error())
	} else {
		fmt.Printf("read user info: %v\n", user)
	}

	//Delete
	delete(db)
}
