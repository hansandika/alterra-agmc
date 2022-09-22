package main

import (
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/internal/http"
	"github.com/hansandika/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	godotenv.Load()
	f := factory.NewFactory()
	e := echo.New()
	middleware.LogMiddleware(e)
	http.NewHttp(e, f)
	e.Logger.Fatal(e.Start(":8080"))
}
