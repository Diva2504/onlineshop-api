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

func CreateTransaction(input *models.TransactionHistory, db *gorm.DB) (models.TransactionHistory, error) {
	var transaction models.TransactionHistory
	result := db.Debug().Create(&input)

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
