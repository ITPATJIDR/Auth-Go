package db 

import (
	"go-jwt-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() () {

	dsn := "root@tcp(127.0.0.1:3306)/gojwt?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
  
  DB = db
	DB.AutoMigrate(&model.User{})
}
