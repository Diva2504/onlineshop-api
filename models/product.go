package models

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Price      int    `gorm:"not null" validate:"required,numeric,min=0,max=50000000"`
	Stock      int    `gorm:"not null" validate:"required,numeric,min=5"`
  CategoryID uint   `json:"category_id"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
  validate := validator.New()
  err := validate.Struct(p)
  if err != nil {
    return err
  }
  return nil
}
