package controllers

import (
	"fmt"
	"net/http"
	. "soft/helper"
	. "soft/models"
	"strconv"
	"strings"

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

func GetHead(c *gin.Context) {
	param := c.Query("id")
	var user_id int
	if len(param) > 10 {
		err, userdata := ParseToken(param)
		if err != 0 {
			c.JSON(http.StatusOK, gin.H{
				"ErrorCode": err,
			})
			return
		}
		user_id = int(userdata["userid"].(float64))
	} else {
		user_id, _ = strconv.Atoi(param)
	}
	var query_ User
	q := DB.Model(&User{}).Where("id = ?", user_id).Find(&query_)
	if q == nil {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 42001,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 42001,
		"Head":      query_.Head,
	})
	return
}

func GetUserInfo(c *gin.Context) {
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

	var res struct {
		Nickname string
		Phone    string
		Gender   string
		Sno      string
		Address  string
		Head     string
	}

	var u User
	DB.Model(&User{}).Where("id = ?", user_id).First(&u)

	res.Address = u.Address
	res.Phone = u.Phone
	res.Gender = u.Sex
	res.Sno = u.Sno
	res.Nickname = u.Nickname
	res.Head = u.Head

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
}

func UpdateUserInfo(c *gin.Context) {
	var post_data struct {
		Token    string
		Nickname string
		Phone    string
		Major    string
		Address  string
		Sno      string
		Gender   string
		Head     string
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

	var u User
	DB.Model(&User{}).Where("id = ?", user_id).First(&u)
	u.Address = post_data.Address
	u.Nickname = post_data.Nickname
	u.Phone = post_data.Phone
	u.Sex = post_data.Gender
	u.Sno = post_data.Sno
	u.Head = post_data.Head
	DB.Save(&u)
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
	res := make([]address_data, 0)
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

func GetDelaultAddress(c *gin.Context) {
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
	res := make([]address_data, 0)
	DB.Model(&Address{}).Where("user_id = ?", user_id).Where("is_default = 1").Find(&address)
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

func UpdateAddress(c *gin.Context) {
	var post_data struct {
		Token         string
		AddressId     int
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
	var address Address
	DB.Model(&Address{}).Where("user_id = ?", user_id).Where("id = ?", post_data.AddressId).First(&address)
	address.DetailAddress = post_data.DetailAddress
	address.IsDefault = post_data.IsDefault
	address.Major = post_data.Major
	address.Name = post_data.Name
	address.Phone = post_data.Phone
	DB.Save(&address)
	if post_data.IsDefault == 1 {
		fmt.Println(address.Id)
		DB.Model(&Address{}).Where("user_id = ?", user_id).Where("is_default = ?", 1).Not("id = ?", address.Id).Update("is_default", 0)
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
	})
}

func GetUserPurchaseGoods(c *gin.Context) {
	token := c.Query("token")
	err, token_data := ParseToken(token)
	if err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}
	user_id := int(token_data["userid"].(float64))
	type a_goods_m struct {
		Title    string
		Price    float32
		Pic      string
		Id       int
		Status   int
		SellerId int
		Time     string
	}
	res := make([]a_goods_m, 0)
	var user_goods []Goods
	DB.Model(&Goods{}).Where("buy_id = ?", user_id).Find(&user_goods)
	for _, c := range user_goods {
		var a_goods a_goods_m
		a_goods.Id = c.Id
		a_goods.Price = c.Price
		a_goods.Title = c.Title
		a_goods.Pic = strings.Split(c.Pics, " ")[0]
		a_goods.Status = c.Status
		a_goods.Time = TimeStampToDate(c.UploadTime)
		a_goods.SellerId = c.UserId
		res = append(res, a_goods)
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
	return
}

func GetUserSaleGoods(c *gin.Context) {
	token := c.Query("token")
	err, token_data := ParseToken(token)
	if err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}
	user_id := int(token_data["userid"].(float64))
	type a_goods_m struct {
		Title   string
		Price   float32
		Pic     string
		Id      int
		Status  int
		Time    string
		BuyerId int
	}
	res := make([]a_goods_m, 0)
	var user_goods []Goods
	DB.Model(&Goods{}).Where("user_id = ?", user_id).Find(&user_goods)
	for _, c := range user_goods {
		var a_goods a_goods_m
		a_goods.Id = c.Id
		a_goods.Price = c.Price
		a_goods.Title = c.Title
		a_goods.Pic = strings.Split(c.Pics, " ")[0]
		a_goods.Status = c.Status
		a_goods.Time = TimeStampToDate(c.UploadTime)
		if c.Status == 0 {
			a_goods.BuyerId = -1
		} else {
			a_goods.BuyerId = c.BuyerId
		}
		res = append(res, a_goods)
	}

	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
	return
}
