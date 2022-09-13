package usecase

import (
	"github.com/hansandika/pkg/users/dto"
	"github.com/hansandika/pkg/users/model"
	"github.com/hansandika/pkg/users/repository"
)

type UsecaseInterface interface {
	CreateNewUser(input *dto.NewUser) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(id int, input *dto.NewUser) (*model.User, error)
	DeleteUser(id int) (*model.User, error)
}

type usecase struct {
	repository repository.RepositoryInterface
}

func InitUsecase(repository repository.RepositoryInterface) UsecaseInterface {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateNewUser(input *dto.NewUser) (*model.User, error) {
	return u.repository.CreateNewUser(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
}

func (u *usecase) GetUserById(id int) (*model.User, error) {
	return u.repository.GetUserById(id)
}

func (u *usecase) GetAllUsers() ([]model.User, error) {
	return u.repository.GetAllUsers()
}

func (u *usecase) UpdateUser(id int, input *dto.NewUser) (*model.User, error) {
	user, err := u.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	return u.repository.UpdateUser(user)
}

func (u *usecase) DeleteUser(id int) (*model.User, error) {
	user, err := u.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	err = u.repository.DeleteUser(user)
	return user, err
}
