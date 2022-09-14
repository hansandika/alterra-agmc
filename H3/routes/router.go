package routes

import (
	"os"

	jwtM "github.com/hansandika/middleware"
	book "github.com/hansandika/pkg/books"
	user "github.com/hansandika/pkg/users"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	userGroup.POST("/login", userController.LoginUser)
	userGroup.POST("", userController.CreateNewUser)

	// middleware
	r := userGroup.Group("/jwt")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.Use(jwtM.HandleAuthJwt)
	r.GET("/:id", userController.GetUserById)
	r.PUT("/:id", userController.UpdateUser)
	r.DELETE("/:id", userController.DeleteUser)

	return e
}
