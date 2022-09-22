package book

import (
	"github.com/labstack/echo"
)

func (c *controller) Route(e *echo.Group) {
	e.GET("", c.GetAllBooks)
	e.POST("", c.CreateNewBook)
	e.GET("/:id", c.GetBookById)
	e.PUT("/:id", c.UpdateBookById)
	e.DELETE("/:id", c.DeleteBookById)
}
