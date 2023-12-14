package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open("notes:tmp_pwd@tcp(127.0.0.1:3306)/notes?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic("failed to connected to database!")
	}

	DB = database
}

func DbMigrate() {
	DB.AutoMigrate(&Note{})
}
