package common

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TableStudent struct {
	Student
	Extra
}

type Student struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
}

type Extra struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (TableStudent) TableName() string {
	return "Students"
}

func (self *Student) Read(db *gorm.DB) error {
	return db.Where(self).First(self).Error
}

func (self *Student) Create(db *gorm.DB) error {
	return db.Create(self).Error
}

func (self *Student) Update(db *gorm.DB) error {
	return db.Save(self).Error
}
