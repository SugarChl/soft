package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Message struct {
	Id      int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	MsgFrom int    `gorm:"type:int(10)"`
	MsgTo   int    `gorm:"type:int(10)"`
	Time    int    `gorm:"type:int(10)"`
	Text    string `gorm:"type:char(200)"`
	Viewed  int    `gorm:"type:int(10);DEFAULT:0"`
	ChatId  int    `gorm:"type:int(10)"`
}

func (a Message) Insert() error {
	return DB.Create(&a).Error
}
