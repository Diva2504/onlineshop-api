package models

import "gorm.io/gorm"

type TransactionHistory struct {
	gorm.Model
	ProductId  int
	UserId     int
	Quantity   int `gorm:"not null" validate:"required"`
	TotalPrice int `gorm:"not null" validate:"required"`
	Product    *Product
	User       *User
}
