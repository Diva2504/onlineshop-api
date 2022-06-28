package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/middleware"
	"github.com/takadev15/onlineshop-api/models"
	"github.com/takadev15/onlineshop-api/repository"
)

type BalanceData struct {
	Balance string `json:"balance"`
}

func (db Handlers) UserRegister(c *gin.Context) {
	var data models.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	res, err := repository.CreateUser(&data, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot register user : %s", err),
		})
	}
	token, err := middleware.GenerateToken(res.ID, res.Email)
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration Succeed",
		"data":    res,
		"token":   token,
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
		"message": "Login Succeed",
		"token":   token,
	})
	return
}

func (db Handlers) UserTopup(c *gin.Context) {
	var data BalanceData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

  balanceData, _ := strconv.Atoi(data.Balance)

	userData := c.MustGet("id")
	userId := uint(userData.(float64))

	res, err := repository.UserTopUp(balanceData, userId, db.Connect)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "message" : err.Error(),
    })
  }
  mess := fmt.Sprintf("updated balanced %d", res)
	c.JSON(http.StatusOK, gin.H{
    "message" : mess,
  })
}

