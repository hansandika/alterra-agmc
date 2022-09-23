package database

import (
	_ "database/sql"
	"fmt"

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
		fmt.Println(url)
		fmt.Println("Ayam")
		panic(err)
	}
	fmt.Println("Connected to DB")
	return db
}

func GetConnection() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}
