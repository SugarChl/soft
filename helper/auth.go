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
	info["AdminID"] = user.Id

	dataByte, _ := json.Marshal(info)
	var dataStr = string(dataByte)
	fmt.Println(dataStr)
	data := jwt.StandardClaims{Subject: dataStr, ExpiresAt: time.Now().Add(time.Hour * time.Duration(720)).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, _ := token.SignedString([]byte(Secret))
	return tokenString
}
