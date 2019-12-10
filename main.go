package main

import (
	"fmt"
	. "soft/controllers"
	. "soft/models"

	"github.com/gin-gonic/gin"

	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	route := gin.Default()
	route.POST("/api/login", Login)
	route.POST("/api/register", Register)
	route.POST("/api/user/info/update", UpdateUserInfo)
	route.POST("/api/user/info/get", GetUserInfo)
	route.GET("/api/user/head/:id", GetHead)

	route.POST("/api/address/new", NewAddress)
	route.POST("/api/address/get", GetAddress)
	route.POST("/api/address/delete", DeleteAddress)
	route.POST("/api/address/update", UpdateAddress)

	route.POST("/api/goods/new", NewGoods)
	route.GET("/api/good/:goodid", GetGood)

	route.GET("/api/community/get", CommunityGetGoods)

	route.POST("/api/file/upload", ImageUpload)
	route.GET("/api/file/:fileid", GetImage)

	route.GET("/api/msglist", GetMsglist)
	route.POST("/api/chat/new", NewMsglist)
	route.POST("/api/message", SendMessage)
	route.GET("/api/message", GetMessage)
	route.Run(":8090")

}
