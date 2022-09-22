package database

import (
	_ "database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func initDB() *gorm.DB {
	url := BaseConfig()
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return db
}

func GetConnection() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}
