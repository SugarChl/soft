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

func NewAddress(c *gin.Context) {
	var post_data struct {
		Token         string
		Name          string
		Phone         string
		Major         string
		DetailAddress string
		IsDefault     int
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
	new_address := Address{
		Name:          post_data.Name,
		IsDefault:     post_data.IsDefault,
		UserId:        user_id,
		Phone:         post_data.Phone,
		DetailAddress: post_data.DetailAddress,
		Major:         post_data.Major,
	}
	if err := new_address.Insert(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	DB.Find(&new_address)
	if post_data.IsDefault == 1 {
		fmt.Println(new_address.Id)
		DB.Model(&Address{}).Where("user_id = ?", user_id).Where("is_default = ?", 1).Not("id = ?", new_address.Id).Update("is_default", 0)
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
	})
}

func GetAddress(c *gin.Context) {
	var post_data struct {
		Token string
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
	var address []Address
	type address_data struct {
		Id            int
		Name          string
		DetailAddress string
		Phone         string
		IsDefault     int
		Major         string
	}
	var res []address_data
	DB.Model(&Address{}).Where("user_id = ?", user_id).Find(&address)
	for _, c := range address {
		var a_data address_data
		a_data.Name = c.Name
		a_data.Phone = c.Phone
		a_data.Id = c.Id
		a_data.DetailAddress = c.DetailAddress
		a_data.IsDefault = c.IsDefault
		a_data.Major = c.Major
		res = append(res, a_data)
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"data":      res,
	})
	return
}

func DeleteAddress(c *gin.Context) {
	var post_data struct {
		Token     string
		AddressId string
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
	var address Address
	q_address := DB.Model(&Address{}).Where("user_id = ?", user_id).Where("id = ?", post_data.AddressId).First(&address)

	if q_address.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 42001,
		})
		return
	}
	if err := address.Delete(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40001,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
	})
}
