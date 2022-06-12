package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/takadev15/onlineshop-api/utils"
)

var secret = "abcd"

func GenerateToken(id uint, email string) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)

  claims["id"] = id
  claims["email"] = email
  claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

  signedToken, err := token.SignedString([]byte(secret)) 
  if err != nil {
    return "", err
  }
  return signedToken, nil
}

func Authentication(c *gin.Context) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    verifiedToken, err := utils.VerifyToken(c)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "err" : "unauthorized",
        "message" : err.Error(),
      })
    }
    c.Set("user_data", verifiedToken)
    c.Next()
  }
}

