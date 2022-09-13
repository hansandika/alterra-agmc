package routes

import (
	book "github.com/hansandika/pkg/books"
	user "github.com/hansandika/pkg/users"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func InitRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	apiv1Group := e.Group("/api/v1")

	// book routes
	bookGroup := apiv1Group.Group("/books")
	bookController := book.InitHttpBookController(db)
	bookGroup.GET("", bookController.GetAllBooks)
	bookGroup.GET("/:id", bookController.GetBookById)
	bookGroup.POST("", bookController.CreateNewBook)
	bookGroup.PUT("/:id", bookController.UpdateBook)
	bookGroup.DELETE("/:id", bookController.DeleteBook)

	// user routes
	userGroup := apiv1Group.Group("/users")
	userController := user.InitHttpUserController(db)
	userGroup.GET("", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetUserById)
	userGroup.POST("", userController.CreateNewUser)
	userGroup.PUT("/:id", userController.UpdateUser)
	userGroup.DELETE("/:id", userController.DeleteUser)

	return e
}
