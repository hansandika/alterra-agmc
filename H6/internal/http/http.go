package http

import (
	"github.com/go-playground/validator"
	"github.com/hansandika/internal/app/auth"
	"github.com/hansandika/internal/app/book"
	"github.com/hansandika/internal/app/user"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/pkg/util"
	"github.com/labstack/echo"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	v1 := e.Group("/api/v1")

	user.NewController(f).Route(v1.Group("/users"))
	book.NewController(f).Route(v1.Group("/books"))
	auth.NewController(f).Route(v1.Group("/auth"))
}
