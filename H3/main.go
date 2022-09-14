package main

import (
	"github.com/go-playground/validator"
	"github.com/hansandika/config"
	"github.com/hansandika/middleware"
	"github.com/hansandika/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.InitDB()
	e := routes.InitRoute(db)
	e.Validator = &middleware.CustomValidator{Validator: validator.New()}
	middleware.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
