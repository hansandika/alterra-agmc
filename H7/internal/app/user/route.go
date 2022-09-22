package user

import (
	"os"

	jwtMiddleware "github.com/hansandika/internal/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (c *controller) Route(e *echo.Group) {

	e.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	e.GET("", c.GetAllUsers)

	r := e.Group("/jwt")
	r.Use(jwtMiddleware.HandleAuthJwt)
	r.GET("/:id", c.GetUserById)
	r.PUT("/:id", c.UpdateUserById)
	r.DELETE("/:id", c.DeleteUserById)
}
