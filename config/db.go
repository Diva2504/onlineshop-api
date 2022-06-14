package config

import (
	"fmt"
	"log"

	"github.com/takadev15/onlineshop-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = "diva"
	dbName   = "onlineshop_db"
	dbPass   = "1234"
	dbPort   = "5432"
	dbHost   = "localhost"
	db       *gorm.DB
	err      error
)

func DBInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, userName, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Databases Error", err.Error())
	}
	log.Printf("Databases Connected")
	db.Debug().AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}
