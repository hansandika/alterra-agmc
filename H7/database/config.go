package database

import "github.com/hansandika/pkg/util"

func BaseConfig() string {
	dbHost := util.Getenv("DB_HOST", "127.0.0.1")
	dbPort := util.Getenv("DB_PORT", "3306")
	dbUser := util.Getenv("DB_USERNAME", "root")
	dbPass := util.Getenv("DB_PASSWORD", "hansgeovani2")
	dbName := util.Getenv("DB_NAME", "altera_tugas_h2")

	return "" +
		dbUser + ":" +
		dbPass + "@(" +
		dbHost + ":" +
		dbPort + ")/" +
		dbName + "?charset=utf8&parseTime=True&loc=Local"
}
