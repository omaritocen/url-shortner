package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Database() *gorm.DB {
	dsn := "root:Jdv%N@fqcV&f8W5%@tcp(127.0.0.1:3306)/url_shortner?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	if err = db.AutoMigrate(&Compression{}); err != nil {
		log.Println(err)
	}

	return db
}
