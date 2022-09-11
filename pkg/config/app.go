package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() { // helps to open  a connection with database
	d, err := gorm.Open("mysql", "pranav:pqper/simplerest?charset=utf&&parseTime=True")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
