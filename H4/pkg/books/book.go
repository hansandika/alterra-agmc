package book

import (
	"github.com/hansandika/pkg/books/controller"
	"github.com/hansandika/pkg/books/model"
	"github.com/hansandika/pkg/books/repository"
	"github.com/hansandika/pkg/books/usecase"
	"github.com/jinzhu/gorm"
)

func InitHttpBookController(db *gorm.DB) *controller.BookHTTPController {
	initMigrate(db)
	repo := repository.InitRepository(db)
	uc := usecase.InitUsecase(repo)
	controller := controller.InitControllerBook(uc)

	return controller
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Book{})
}
