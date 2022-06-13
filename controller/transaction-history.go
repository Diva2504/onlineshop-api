package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/models"
	"github.com/takadev15/onlineshop-api/repository"
)

type ResforTransaction struct {
	ID         int           `json:"id"`
	ProductId  int           `json:"product_id"`
	UserID     int           `json:"user_id"`
	TotalPrice int           `json:"total_price"`
	Quantity   int           `json:"quantity"`
	Product    ResforProduct `json:"product"`
	User       ResforUser    `json:"user"`
}

type ResforProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId int       `json:"category_id"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}
type ResforUser struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	Balance    int       `json:"balance"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}
type InputTransaction struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (db Handlers) GetforAdmin(c *gin.Context) {
	var transactionRes []ResforTransaction
	res, err := repository.GetforAdmin(db.Connect)

	for i := range res {
		transactionRes[i].ID = int(res[i].ID)
		transactionRes[i].ProductId = res[i].ProductId
		transactionRes[i].UserID = res[i].UserId
		transactionRes[i].TotalPrice = res[i].TotalPrice
		transactionRes[i].Quantity = res[i].Quantity
		transactionRes[i].Product.ID = int(res[i].Product.ID)

		transactionRes[i].Product.Title = res[i].Product.Title
		transactionRes[i].Product.Price = res[i].Product.Price
		transactionRes[i].Product.Stock = res[i].Product.Stock
		transactionRes[i].Product.CategoryId = res[i].Product.CategoryId
		transactionRes[i].Product.Created_at = res[i].Product.CreatedAt
		transactionRes[i].Product.Updated_at = res[i].Product.UpdatedAt

		transactionRes[i].User.ID = int(res[i].User.ID)
		transactionRes[i].User.Email = res[i].User.Email
		transactionRes[i].User.FullName = res[i].User.FullName
		transactionRes[i].User.Balance = res[i].User.Balance
		transactionRes[i].User.Created_at = res[i].User.CreatedAt
		transactionRes[i].User.Updated_at = res[i].User.UpdatedAt
	}

	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"transaction": transactionRes,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) GetforUser(c *gin.Context) {
	var transactionRes []ResforTransaction
	res, err := repository.GetforAdmin(db.Connect)

	for i := range res {
		transactionRes[i].ID = int(res[i].ID)
		transactionRes[i].ProductId = res[i].ProductId
		transactionRes[i].UserID = res[i].UserId
		transactionRes[i].TotalPrice = res[i].TotalPrice
		transactionRes[i].Quantity = res[i].Quantity
		transactionRes[i].Product.ID = int(res[i].Product.ID)

		transactionRes[i].Product.Title = res[i].Product.Title
		transactionRes[i].Product.Price = res[i].Product.Price
		transactionRes[i].Product.Stock = res[i].Product.Stock
		transactionRes[i].Product.CategoryId = res[i].Product.CategoryId
		transactionRes[i].Product.Created_at = res[i].Product.CreatedAt
		transactionRes[i].Product.Updated_at = res[i].Product.UpdatedAt
	}

	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"my_transaction": transactionRes,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) MakeTransaction(c *gin.Context) {
	var (
		product models.Product
		input   InputTransaction
	)

}
