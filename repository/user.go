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

func UserTopUp(data int, userId uint, db *gorm.DB) (int, error){
  var user models.User
  if err := db.Debug().First(&user, userId).Error; err != nil {
    return 0, err
  }
  err := db.Debug().Model(&user).Update("balance", data).Error
  if err != nil {
    return 0, err
  }
  return data, nil
}
