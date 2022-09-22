package auth

import (
	"github.com/labstack/echo"
)

func (c *controller) Route(e *echo.Group) {
	e.POST("/login", c.LoginByEmailAndPassword)
	e.POST("/signup", c.RegisterUserByEmailAndPassword)
}
