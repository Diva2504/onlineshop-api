package repository

import (
	"github.com/takadev15/onlineshop-api/models"
	"gorm.io/gorm"
)

func GetforAdmin(db *gorm.DB) ([]models.TransactionHistory, error) {
	var transaction []models.TransactionHistory
	result := db.Find(&transaction)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return transaction, result.Error
		}
	}
}

func CreateTransaction(productId uint, quantity int, userId uint, db *gorm.DB) (models.TransactionHistory, error) {
	var transaction models.TransactionHistory
	var product models.Product
	var user models.User

	err := db.Where("id = ?", productId).Take(&product).Error

	if err != nil && product.Stock < quantity {
		return models.TransactionHistory{}, err
	}
	totalPrice := quantity * product.Price

	err = db.Where("id = ?", userId).Take(&user).Error
	if err != nil && user.Balance < totalPrice {
		return models.TransactionHistory{}, err
	}

	transaction.ProductId = productId
	transaction.Quantity = quantity
	transaction.UserId = userId
	transaction.TotalPrice = totalPrice
	transaction.Product = &product
	transaction.User = &user

	result := db.Debug().Create(&transaction)

	if result.Error != nil {
		return models.TransactionHistory{}, result.Error
	}
	return transaction, nil
}

func GetforUser(db *gorm.DB) ([]models.TransactionHistory, error) {
	var transaction []models.TransactionHistory
	result := db.Find(&transaction)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return transaction, result.Error
		}
	}
}
