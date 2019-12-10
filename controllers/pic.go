package controllers

import (
	"net/http"

	"fmt"
	"os"
	. "soft/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UUID() string {
	Id := uuid.New()
	return Id.String()
}

func GetImage(c *gin.Context) {
	fileid := c.Param("fileid")
	_, err := os.Stat("static/" + fileid)

	if err == nil {
		c.File("static/" + fileid)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 42001,
	})
	return
}

func ImageUpload(c *gin.Context) {
	Token := c.Query("token")

	if err, _ := ParseToken(Token); err != 0 {
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": err,
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ErrorCode": 40002,
		})
		return
	}
	filename := file.Filename

	new_name := UUID()
	index := strings.Index(filename, ".")
	if index == -1 {
		filename = new_name + ".jpg"
	} else {
		filename = strings.Replace(filename, filename[:index], new_name, 1)
	}
	c.SaveUploadedFile(file, "static/"+filename)
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": 0,
		"FileUrl":   "http://soft.sugarchl.top/api/file/" + filename,
	})
	return
}
