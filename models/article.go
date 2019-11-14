package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Article struct {
	Id               int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserId           int    `gorm:"type:int(10)"`
	Title            string `gorm:"type:char(200)"`
	Content          string `gorm:"type:char(200)"`
	UploadTime       int    `gorm:"type:int(10)"`
	CommentNumber    int    `gorm:"type:int(10);DEFAULT:0"`
	LikeNumber       int    `gorm:"type:int(10);DEFAULT:0"`
	CollectionNumber int    `gorm:"type:int(10);DEFAULT:0"`
}

func (a Article) Insert() error {
	return DB.Create(&a).Error
}
