package main

import (
	"database/sql"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/* gorm 支持的 struct tag
Column	        指定列名
Type	        指定列数据类型
Size	        指定列大小, 默认值255
PRIMARY_KEY	    将列指定为主键
UNIQUE	        将列指定为唯一
DEFAULT	        指定列默认值
PRECISION	    指定列精度
NOT NULL	    将列指定为非 NULL
AUTO_INCREMENT	指定列是否为自增类型
INDEX	        创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
UNIQUE_INDEX	和 INDEX 类似，只不过创建的是唯一索引
EMBEDDED	    将结构设置为嵌入
EMBEDDED_PREFIX	设置嵌入结构的前缀
-	            忽略此字段

用于表关联相关的 struct tag

MANY2MANY	                        指定连接表
FOREIGNKEY	                        设置外键
ASSOCIATION_FOREIGNKEY	            设置关联外键
POLYMORPHIC	                        指定多态类型
POLYMORPHIC_VALUE	                指定多态值
JOINTABLE_FOREIGNKEY	            指定连接表的外键
ASSOCIATION_JOINTABLE_FOREIGNKEY	指定连接表的关联外键
SAVE_ASSOCIATIONS	                是否自动完成 save 的相关操作
ASSOCIATION_AUTOUPDATE	            是否自动完成 update 的相关操作
ASSOCIATION_AUTOCREATE	            是否自动完成 create 的相关操作
ASSOCIATION_SAVE_REFERENCE	        是否自动完成引用的 save 的相关操作
PRELOAD	                            是否自动完成预加载的相关操作
*/

type User struct {
	gorm.Model //gorm 内置的惯例字段，可用可不用
	// CreatedAt 自动设置为记录的首次时间：db.Create(&user)
	// 如果更新用 Update 方法: db.Model(&user).Update("CreatedAt", time.Now())
	// UpdatedAt 在记录变更时自动更新：db.Save(&user)
	// DeletedAt 在删除时设置为删除时间（如果有 DeleteAt 字段，默认使用软删除）
	Name         string
	Ages         sql.NullInt32
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"` //类型，唯一索引
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"` //忽略字段
}

// 可以用此方法修改表名，也可以在创建表的时候修改
func (User) TableName() string {
	return "users"
}

/* 根据情况，返回不同的表名
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}
*/

func main() {
	db, err := gorm.Open("mysql", "gorm:passwd123@(localhost)/gormdb?charset=utf8&parseTime=True&loc=Local")
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

	if !db.HasTable(&User{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	}
}
