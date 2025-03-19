package controllers

import (
	"log"

	"github.com/Wasay1567/url-shortner-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:Wasay786@@tcp(127.0.0.1:3306)/urls?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.AutoMigrate(&models.URL{})
}
