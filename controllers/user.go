package controllers

import (
	"fmt"
	"net/http"
	. "soft/helper"
	. "soft/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var post_data struct {
		Username string
		Passwd   string
	}
	if err := c.BindJSON(&post_data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40002,
		})
		return
	}

	var query_ User
	q_user := DB.Model(&User{}).Where("username = ?", post_data.Username).First(&query_)
	if q_user.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 41001,
		})
		return
	} else if query_.Passwd != post_data.Passwd {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 41002,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Token":     GenerateToken(query_),
	})
}

func Register(c *gin.Context) {
	var post_data struct {
		Username string
		Passwd   string
		Sno      string
		Spasswd  string
	}
	if err := c.BindJSON(&post_data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40002,
		})
		return
	}
	if len(post_data.Passwd) < 5 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 41003,
		})
		return
	}

	var query_ User
	q_user := DB.Model(&User{}).Where("username = ?", post_data.Username).First(&query_)
	if q_user.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 41003,
		})
		return
	}
	new_user := User{
		Username: post_data.Username,
		Passwd:   post_data.Passwd,
		Sno:      post_data.Sno,
		Spwd:     post_data.Spasswd,
	}
	if err := new_user.Insert(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
	})

}
