package controllers

import (
	"fmt"
	"net/http"
	. "soft/helper"
	. "soft/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NewGoods(c *gin.Context) {
	var post_data struct {
		Token   string
		Title   string
		Content string
		Price   float32
		Pics    []string
	}
	if err := c.BindJSON(&post_data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40002,
		})
		return
	}
	err, token_data := ParseToken(post_data.Token)
	if err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}

	user_id := int(token_data["userid"].(float64))
	new_goods := Goods{
		Title:      post_data.Title,
		Content:    post_data.Content,
		UserId:     user_id,
		UploadTime: int(time.Now().Unix()),
		Price:      post_data.Price,
		Pics:       strings.Join(post_data.Pics, " "),
	}
	if err := new_goods.Insert(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	DB.Find(&new_goods)
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"GoodsId":   new_goods.Id,
	})
}
