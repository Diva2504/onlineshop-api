package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = "abcdefghijklmnopq"

func VerifyToken(c *gin.Context) (interface{}, error) {
	errCode := errors.New("please use correct authentication")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
    log.Printf("Error 1 : %s", errCode)
		return nil, errCode
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      log.Printf("Error 2 : %s", errCode)
			return nil, errCode
		}
		return []byte(secret), nil
	})
  if err != nil {
    log.Printf("Error 3 : %s", err)
    return nil, err
  }

	if !token.Valid {
    log.Printf("Error 4 : %s", errCode)
		return nil, errCode
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
    log.Printf("Error 5 : %s", errCode)
		return nil, errCode
	}

	return token.Claims.(jwt.MapClaims), nil
}
