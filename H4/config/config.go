package config

import (
	"os"

	_ "database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func BaseConfig() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbUser == "" {
		dbUser = "root"
	}
	if dbPass == "" {
		dbPass = "hansgeovani2"
	}
	if dbName == "" {
		dbName = "altera_tugas_h2"
	}

	return "" +
		dbUser + ":" +
		dbPass + "@(" +
		dbHost + ":" +
		dbPort + ")/" +
		dbName + "?charset=utf8&parseTime=True&loc=Local"
}

func InitDB() *gorm.DB {
	url := BaseConfig()
	db, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return db
}
