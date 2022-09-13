package main

import (
	"github.com/hansandika/config"
	"github.com/hansandika/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.InitDB()
	e := routes.InitRoute(db)

	e.Logger.Fatal(e.Start(":8080"))
}
