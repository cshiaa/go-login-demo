package utils

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

)


func ValidToken(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("xxxx"), nil
	})
	if err!= nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Request.Header.Get("atoken")
    if token != "" {
		return token
    }
	beareToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(beareToken, " ")) == 2 {
		return strings.Split(beareToken, " ")[1]
	}
    return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	user, err := VerifyToken(tokenString)
	if err!= nil {
        return 0, err
    }
	user_id := user.UserID
	return user_id, nil
}