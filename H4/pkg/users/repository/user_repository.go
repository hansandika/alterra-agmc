package repository

import (
	"github.com/hansandika/pkg/users/model"
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	CreateNewUser(user *model.User) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNewUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	user, err = r.GetUserById(int(user.ID))
	return user, err
}

func (r *repository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Find(&user, id).Error
	return &user, err
}

func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return &user, err
}

func (r *repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(user *model.User) error {
	err := r.db.Delete(&user).Error
	return err
}
