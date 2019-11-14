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
	route.Run(":8090")

}
