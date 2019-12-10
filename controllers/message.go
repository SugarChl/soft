package controllers

import (
	"fmt"
	"net/http"
	. "soft/helper"
	. "soft/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var post_data struct {
		Token  string
		ChatId int
		Text   string
		To     int
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
	new_msg := Message{
		Text:    post_data.Text,
		MsgFrom: user_id,
		MsgTo:   post_data.To,
		Time:    int(time.Now().Unix()),
		ChatId:  post_data.ChatId,
	}
	if err := new_msg.Insert(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	DB.Find(&new_msg)
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
	})
}

func GetMessage(c *gin.Context) {
	type a_msg_block struct {
		Time string
		Text string
		Mine int
	}
	token := c.Query("token")
	chat := c.Query("chat_id")
	err, token_data := ParseToken(token)
	if err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}
	user_id := int(token_data["userid"].(float64))
	chat_id, errt := strconv.Atoi(chat)
	if errt != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40002,
		})
		return
	}
	msgs := make([]a_msg_block, 0)
	var query_ []Message
	DB.Model(&Message{}).Where("chat_id = ?", chat_id).Find(&query_)
	for _, c := range query_ {
		var a_msg a_msg_block
		a_msg.Text = c.Text
		a_msg.Time = TimeStampToDate(c.Time)
		if c.MsgFrom == user_id {
			a_msg.Mine = 1
		}
		msgs = append(msgs, a_msg)
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      msgs,
	})
	return

}
