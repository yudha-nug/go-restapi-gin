package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:skywalker@tcp(192.168.1.6:3306)/go_restapi_gin"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})
	DB = database
}
