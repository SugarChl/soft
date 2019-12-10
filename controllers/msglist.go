package controllers

import (
	"fmt"
	"net/http"
	. "soft/helper"
	. "soft/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMsglist(c *gin.Context) {
	token := c.Query("token")
	err, token_data := ParseToken(token)
	if err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}
	type a_msglist_block struct {
		ChatId      int
		MsgFromId   int
		MsgFromName string
		MsgFromHead string
		LatestText  string
		LatestTime  string
	}
	user_id := int(token_data["userid"].(float64))

	res := make([]a_msglist_block, 0)
	var query_ []MsgList
	DB.Model(&MsgList{}).Where("msg_to = ?", user_id).Or("msg_from = ?", user_id).Find(&query_)
	for _, c := range query_ {
		var a_msglist_b a_msglist_block
		a_msglist_b.ChatId = c.ChatId

		var a_msg Message
		msg_in := DB.Model(&Message{}).Where("chat_id = ?", c.ChatId).Order("Time").First(&a_msg).Error
		if msg_in == nil {
			a_msglist_b.LatestText = a_msg.Text
			a_msglist_b.LatestTime = strconv.Itoa(a_msg.Time)
		}

		var a_user User
		var q_id int
		if c.MsgFrom == user_id {
			q_id = c.MsgTo
		} else {
			q_id = c.MsgFrom
		}
		DB.Model(&User{}).Where("id = ?", q_id).First(&a_user)
		a_msglist_b.MsgFromHead = a_user.Head
		a_msglist_b.MsgFromName = a_user.Nickname
		a_msglist_b.MsgFromId = q_id
		res = append(res, a_msglist_b)
	}
	var temp a_msglist_block
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i].LatestTime < res[j].LatestTime {
				temp = res[i]
				res[i] = res[j]
				res[j] = temp
			}
		}
	}
	for i := 0; i < len(res); i++ {
		temp_t := res[i].LatestTime
		int_temp_t, _ := strconv.Atoi(temp_t)
		res[i].LatestTime = TimeStampToDate(int_temp_t)
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
	return
}

func NewMsglist(c *gin.Context) {
	var post_data struct {
		Token string
		MsgTo int
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

	var msgtoUser User
	user_in_db := DB.Model(&User{}).Where("user_id = ?", post_data.MsgTo).Find(&msgtoUser).Error
	if user_in_db != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 42001,
		})
		return
	}

	new_msglist := MsgList{
		MsgTo:   post_data.MsgTo,
		MsgFrom: user_id,
	}
	if err := new_msglist.Insert(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	DB.Find(&new_msglist)
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"ChatId":    new_msglist.ChatId,
	})
}
