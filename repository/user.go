package repository

import (

	"github.com/takadev15/onlineshop-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(data *models.User, db *gorm.DB) (models.User, error) {
  var user models.User
  err := db.Debug().Create(&data).Error
  if err != nil {
    return models.User{}, err
  }
  user = *data
  return user, nil
}

func UserLogin(data *models.User, db *gorm.DB) (models.User, error) {
  var user models.User
  password := data.Password
  err := db.Debug().Where("email = ?", data.Email).Take(&user).Error
  if err != nil {
    return models.User{}, err
  }
  comparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if comparePass != nil {
    return models.User{}, comparePass
  }
  return user, nil
}
