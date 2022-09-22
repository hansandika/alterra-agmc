package repository

import (
	"github.com/hansandika/internal/model"
	"github.com/jinzhu/gorm"
)

type UserRepositoryInterface interface {
	ValidateUserExists(email string) (bool, error)
	CreateNewUser(user *model.User) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) ValidateUserExists(email string) (bool, error) {
	var count int64

	if err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *userRepository) CreateNewUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	user, err = r.GetUserById(int(user.ID))
	return user, err
}

func (r *userRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.db.Find(&user, id).Error
	return &user, err
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return &user, err
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(user *model.User) error {
	err := r.db.Delete(&user).Error
	return err
}
