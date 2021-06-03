package common

import (
	"acussm/demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {



	dsn := "root:root@tcp(127.0.0.1:3306)/accussm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil {
		panic("failed to connect database"+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return  db
}

func GetDB() *gorm.DB {

	return DB
	
}