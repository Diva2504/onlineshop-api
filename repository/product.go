package repository

import (
	"github.com/takadev15/onlineshop-api/models"
	"gorm.io/gorm"
)

func GetAllProduct(db *gorm.DB) ([]models.Product, error) {
	var product []models.Product
	result := db.Find(&product)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return product, result.Error
		}
	}
}
func GetProduct(id int, db *gorm.DB) (models.Product, error) {
	var product models.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func CreateProduct(req *models.Product, db *gorm.DB) error {
	result := db.Create(&req)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProduct(id int, data *models.Product, db *gorm.DB) (models.Product, error) {
	var product models.Product
	err := db.Model(&product).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, err
}

func DeleteProduct(id int, db *gorm.DB) error {
	var product models.Product

	del := db.Delete(&product, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}
}
