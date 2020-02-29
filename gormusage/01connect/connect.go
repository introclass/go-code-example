package main

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// 准备数据库
	// create database gormdb;
	// create user 'gorm'@'%'  identified by 'passwd123';
	// alter user 'gorm'@'%' identified with mysql_native_password by 'password123';
	// grant all on gormdb.* to 'gorm'@'%';
	// flush privileges;

	//mysql, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
	mysql, err := gorm.Open("mysql", "gorm:passwd123@(localhost)/gormdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		glog.Fatal("mysql: ", err.Error())
	}
	defer mysql.Close()

	sqlite, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		glog.Fatal("sqlite3: ", err.Error())
	}
	defer sqlite.Close()

	// 准备数据库
	// create user gorm with password 'password123';
	// create database gormdb;
	// grant all on database gormdb to gorm;
	//
	// 在 pg_hba.conf 中添加配置：
	// host     gormdb    gorm    0.0.0.0/0    password
	// hostssl  gormdb    gorm    0.0.0.0/0    password
	//
	// 命令行登陆
	// psql -U gorm  gormdb -h 127.0.0.1 -p 5432

	//postgres, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gorm dbname=gormdb password=password123 sslmode=disable")
	postgres, err := gorm.Open("postgres", "postgres://gorm:password123@localhost/gormdb?sslmode=disable")
	if err != nil {
		glog.Fatal("postgres: ", err.Error())
	}
	defer postgres.Close()

	//mssql, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	//if err != nil {
	//	glog.Fatal("mssql: ", err.Error())
	//}
	//defer mssql.Close()

}
