package user

import (
	"github.com/hansandika/pkg/users/controller"
	"github.com/hansandika/pkg/users/model"
	"github.com/hansandika/pkg/users/repository"
	"github.com/hansandika/pkg/users/usecase"
	"github.com/jinzhu/gorm"
)

func InitHttpUserController(db *gorm.DB) *controller.UserHTTPController {
	initMigrate(db)
	repo := repository.InitRepository(db)
	uc := usecase.InitUsecase(repo)
	controller := controller.InitControllerUser(uc)

	return controller
}

func initMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
