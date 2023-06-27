package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"url-shortner/models"
)

var DB *gorm.DB

func Init() {
	dsn := "root:Jdv%N@fqcV&f8W5%@tcp(127.0.0.1:3306)/url_shortner?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	if err = db.AutoMigrate(&models.Compression{}); err != nil {
		log.Println(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
