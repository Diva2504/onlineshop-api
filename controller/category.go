package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/models"
	"github.com/takadev15/onlineshop-api/repository"
	// "github.com/takadev15/onlineshop-api/models"
)

func (db Handlers) CreateCategory(c *gin.Context) {
	var data models.Category
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	res, err := repository.CreateCategory(&data, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": res,
	})
}

func (db Handlers) GetCategory(c *gin.Context) {
	res, err := repository.GetAllCategory(db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (db Handlers) UpdateCategory(c *gin.Context) {

}

func (db Handlers) DeleteCategory(c *gin.Context) {}
