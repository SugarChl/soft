package controllers

import (
	//"fmt"
	"net/http"
	//. "soft/helper"
	. "soft/models"
	//"time"

	"github.com/gin-gonic/gin"
)

func CommunityGetGoods(c *gin.Context) {
	type Res_Goods struct {
		GoodsId    int
		Title      string
		Pics       []string
		Price      float32
		SoldNumber int
		Content    string
	}
	var res []Res_Goods
	var goods []Goods
	DB.Limit(10).Find(&goods)
	for _, c := range goods {
		var a_data Res_Goods
		a_data.GoodsId = c.Id
		a_data.SoldNumber = c.SoldNumber
		a_data.Title = c.Title
		a_data.Price = c.Price
		a_data.Content = c.Content

		res = append(res, a_data)
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
	return
}
