package config

import (
	"os"

	_ "database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func BaseConfig() string {
	return "" +
		os.Getenv("DB_USERNAME") + ":" +
		os.Getenv("DB_PASSWORD") + "@(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
}

func InitDB() *gorm.DB {
	url := BaseConfig()
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return db
}
