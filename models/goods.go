package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Goods struct {
	Id               int     `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserId           int     `gorm:"type:int(10)"`
	Title            string  `gorm:"type:char(200)"`
	Content          string  `gorm:"type:char(200)"`
	UploadTime       int     `gorm:"type:int(10)"`
	LikeNumber       int     `gorm:"type:int(10);DEFAULT:0"`
	CollectionNumber int     `gorm:"type:int(10);DEFAULT:0"`
	Pics             string  `gorm:"type:char(200);DEFAULT:0"`
	Price            float32 `gorm:"type:float"`
	SoldNumber       int     `gorm:"type:int(10);DEFAULT:0"`
}

func (a Goods) Insert() error {
	return DB.Create(&a).Error
}
