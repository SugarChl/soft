package controllers

import (
	//"fmt"
	"net/http"
	//. "soft/helper"
	. "soft/models"
	//"time"

	"github.com/gin-gonic/gin"
)

func CommunityGetArticles(c *gin.Context) {
	type Res_Article struct {
		ArticleId        int
		Title            string
		CommentNumber    int
		LikeNumber       int
		CollectionNumber int
		UserName         string
		UserHead         string
		Content          string
	}
	var res []Res_Article
	var articles []Article
	DB.Limit(5).Find(&articles)
	for _, c := range articles {
		var a_data Res_Article
		a_data.ArticleId = c.Id
		a_data.CollectionNumber = c.CollectionNumber
		a_data.CommentNumber = c.CommentNumber
		a_data.LikeNumber = c.LikeNumber
		a_data.Title = c.Title
		a_data.Content = c.Content

		var article_user User
		DB.Model(&User{}).Where("id = ?", c.UserId).First(&article_user)

		a_data.UserName = article_user.Username
		res = append(res, a_data)
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"Data":      res,
	})
	return
}
