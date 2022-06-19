package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/middleware"
	"github.com/takadev15/onlineshop-api/models"
	"github.com/takadev15/onlineshop-api/repository"
)

func (db Handlers) UserRegister(c *gin.Context) {
  var data models.User
  if err := c.ShouldBindJSON(&data); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
  }
  res , err := repository.CreateUser(&data, db.Connect)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message" : fmt.Sprintf("Cannot register user : %s", err),
    })
  }
  token, err := middleware.GenerateToken(res.ID, res.Email)
  c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
  c.JSON(http.StatusCreated, gin.H{
    "message" : "Registration Succeed",
    "data" : res,
  })
}

func (db Handlers) UserLogin(c *gin.Context) {
  var data models.User
  if err := c.ShouldBindJSON(&data); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
    return
  }
  res, err := repository.UserLogin(&data, db.Connect)
  if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message": fmt.Sprintf("Cannot Log in User : %s", err),
    })
    return
  }
  token, err := middleware.GenerateToken(res.ID, res.Email)
  c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
  c.JSON(http.StatusOK, gin.H{
    "message" : "Login Succeed",
  })
  return
}

func (db Handlers) UserTopup(c *gin.Context) {
  var data int
  if err := c.ShouldBindJSON(&data); err != nil {
    c.AbortWithStatus(http.StatusBadRequest)
  }
  res, err := repository.UserTopUp(data, 1, db.Connect)
  c.JSON(http.StatusOK, gin.H{})
}

