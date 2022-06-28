package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type               string
	SoldProductAmmount int `json:"sold_ammount"`
  Products           []Product
}
