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
	Status           int     `gorm:"type:int(10);DEFAULT:0"` // 0未售出   1已售出  2订单结束
	BuyerId          int     `gorm:"type:int(10);DEFAULT:0"`
	BuyTime          int     `gorm:"type:int(10);DEFAULT:0"`
}

func (a Goods) Insert() error {
	return DB.Create(&a).Error
}

func (a Goods) Delete() error {
	return DB.Delete(&a).Error
}
