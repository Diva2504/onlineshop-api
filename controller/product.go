package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/models"
	"github.com/takadev15/onlineshop-api/repository"
)

type ResponseProduct struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
	CategoryId uint      `json:"category_id"`
}

func (db Handlers) GetAllProduct(c *gin.Context) {
	res, err := repository.GetAllProduct(db.Connect)

	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"products": res,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) GetProduct(c *gin.Context) {
	var (
		result     gin.H
		productres ResponseProduct
	)
	productId, _ := strconv.Atoi(c.Param("id"))
	res, err := repository.GetProduct(productId, db.Connect)
	{
		productres.ID = res.ID
		productres.Title = res.Title
		productres.Price = res.Price
		productres.Stock = res.Stock
		productres.CategoryId = res.CategoryID
		productres.Created_at = res.CreatedAt
	}
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"product": productres,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateProduct(c *gin.Context) {
	var (
		product models.Product
		result  gin.H
	)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateProduct(product, db.Connect)
	if err != nil {
	  c.JSON(http.StatusInternalServerError, gin.H{
      "message" : err,
    })
    return
	}
	result = gin.H{
		"message": "sucess",
    "Added Item" : product.Title,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) DeleteProduct(c *gin.Context) {
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	err := repository.DeleteProduct(id, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "product has been succesfully deleted",
	})
}

func (db Handlers) UpdateProduct(c *gin.Context) {
	var (
		product models.Product
		result  gin.H
	)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	productId, _ := strconv.Atoi(c.Param("id"))
	res, err := repository.UpdateProduct(productId, &product, db.Connect)
	if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "message" : err.Error(),
    })
	}
	result = gin.H{
		"product": res,
	}
	c.JSON(http.StatusCreated, result)
}

