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

var counter uint

func (db Handlers) GetAllProduct(c *gin.Context) {
	var productres []ResponseProduct
	res, err := repository.GetAllProduct(db.Connect)

	for i := range res {
		productres[i].ID = res[i].ID
		productres[i].Title = res[i].Title
		productres[i].Price = res[i].Price
		productres[i].Stock = res[i].Stock
		productres[i].CategoryId = res[i].CategoryId
		productres[i].Created_at = res[i].CreatedAt
	}

	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"products": productres,
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
		productres.CategoryId = res.CategoryId
		productres.Created_at = res.CreatedAt
	}
	if err != nil {
		result = gin.H{
			"message": err,
		}
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
	err := repository.CreateProduct(&product, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"product": product,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) DeleteProduct(c *gin.Context) {
	var result gin.H
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	err := repository.DeleteProduct(id, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"message": "Product has been successfully deleted",
	}
	c.JSON(http.StatusOK, result)
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
	_, err := repository.UpdateProduct(productId, &product, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"product": product,
	}
	c.JSON(http.StatusOK, result)
}
