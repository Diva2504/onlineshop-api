package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title      string `gorm:"not null" validate:"required,title"`
	Price      int    `gorm:"not null" validate:"required,price,min=0,max=50.000.000"`
	Stock      int    `gorm:"not null" validate:"required,stock,min=5"`
	CategoryId int
	Category   *Category
}
