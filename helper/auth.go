package helper

import (
	"encoding/json"
	"fmt"
	. "soft/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const Secret = "Chl Fuzhou University"

func GenerateToken(user User) string {
	info := map[string]interface{}{}
	info["username"] = user.Username
	info["userid"] = user.Id

	dataByte, _ := json.Marshal(info)
	var dataStr = string(dataByte)
	fmt.Println(dataStr)
	data := jwt.StandardClaims{Subject: dataStr, ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, _ := token.SignedString([]byte(Secret))
	return tokenString
}

func ParseToken(token string) (int, map[string]interface{}) {
	var tokendata map[string]interface{}
	tokenInfo, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Secret), nil
	})
	if err != nil {
		fmt.Println(err)
		return 40300, tokendata
	}
	err = tokenInfo.Claims.Valid()
	fmt.Println(err)
	if err != nil {
		return 40300, tokendata //非法token
	}
	finToken := tokenInfo.Claims.(jwt.MapClaims)
	succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	if succ == false {
		return 40400, tokendata //过期token
	}
	data := finToken["sub"]
	json.Unmarshal([]byte(data.(string)), &tokendata)

	return 0, tokendata
}
