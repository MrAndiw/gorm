package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "root:23031996@tcp(127.0.0.1:3306)/optimize_order?charset=utf8mb4&parseTime=True&loc=Local"
	dbcon, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database")
	}

	db = dbcon
}

func GetConnection() *gorm.DB {
	return db
}
