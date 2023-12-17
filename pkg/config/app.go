package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	dsn := "root:govind@123@tcp(localhost:3306)/books?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
