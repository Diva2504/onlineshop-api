package repository

import (
	"github.com/takadev15/onlineshop-api/models"
	"gorm.io/gorm"
)

func CreateCategory(data *models.Category, db *gorm.DB) (models.Category, error){
  err := db.Debug().Create(&data).Error
  if err != nil {
    return models.Category{}, err
  }
  category := *data
  return category, nil
}

func GetAllCategory(db *gorm.DB) ([]models.Category, error){
  var categories []models.Category
  err := db.Find(&categories).Error
  if err != nil {
    return nil, err
  }
  return categories, nil
}

func GetCategory(id int, db *gorm.DB) (models.Category, error){
  var category models.Category
  if err := db.Debug().Where("id = ?", id).First(&category).Error; err != nil {
    return models.Category{}, err
  }
  return category, nil
}

func UpdateCategory(id int, data *models.Category,db *gorm.DB) (models.Category, error){
  var categories models.Category

  err := db.Preload("Items").First(&categories, id).Error
  if err != nil {
    return models.Category{}, err
  }
  err = db.Model(&categories).Updates(&data).Error
  if err != nil {
    return models.Category{}, err
  }
  return categories, nil
}

func DeleteCategory(id int, db *gorm.DB) error{
  var categories models.Category
  err := db.First(&categories, id).Error
  if err != nil {
    return err
  }
  err = db.Delete(&categories).Error
  if err != nil {
    return err
  } else {
    return nil
  }
}
