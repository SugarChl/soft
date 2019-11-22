package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Username string `gorm:"type:char(100)"`
	Passwd   string `gorm:"type:char(100)"`
	Sno      string `gorm:"type:char(100)"`
	Spwd     string `gorm:"type:char(100)"`
	Sex      string `gorm:"type:char(10)"`
	Address  string `gorm:"type:char(100)"`
	Phone    string `gorm:"type:char(20)"`
	Nickname string `gorm:"type:char(50)"`
}

type Address struct {
	Id            int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserId        int    `gorm:"type:int(5)"`
	Name          string `gorm:"type:char(100)"`
	Phone         string `gorm:"type:char(100)"`
	Major         string `gorm:"type:char(100)"`
	DetailAddress string `gorm:"type:char(200)"`
	IsDefault     int    `gorm:"type:int(2)"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:ab123@tcp(localhost)/soft?charset=utf8")
	if err == nil {
		DB = db
		db.SingularTable(true)

		db.AutoMigrate(
			&User{},
			&Address{},
			&Goods{},
		)
		return db, err
	}
	return nil, err
}
func (u User) Insert() error {
	return DB.Create(&u).Error
}

func (u Address) Insert() error {
	return DB.Create(&u).Error
}
func (u Address) Delete() error {
	return DB.Delete(&u).Error
}
