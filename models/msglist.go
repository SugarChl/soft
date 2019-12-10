package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MsgList struct {
	ChatId  int `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	MsgTo   int `gorm:"type:int(10)"`
	MsgFrom int `gorm:"type:int(10)"`
}

func (a MsgList) Insert() error {
	return DB.Create(&a).Error
}
func (u MsgList) Delete() error {
	return DB.Delete(&u).Error
}
