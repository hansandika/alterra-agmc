package usecase

import (
	"time"

	"github.com/hansandika/auth"
	"github.com/hansandika/pkg/users/dto"
	"github.com/hansandika/pkg/users/model"
	"github.com/hansandika/pkg/users/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsecaseInterface interface {
	CreateNewUser(input *dto.NewUser) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserCredential(input *dto.UserCredential) (string, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User, input *dto.NewUser) (*model.User, error)
	DeleteUser(user *model.User) (*model.User, error)
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
	password := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return u.repository.CreateNewUser(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	})
}

func (u *usecase) GetUserById(id int) (*model.User, error) {
	return u.repository.GetUserById(id)
}

func (u *usecase) GetUserByEmail(email string) (*model.User, error) {
	return u.repository.GetUserByEmail(email)
}

func (u *usecase) GetUserCredential(input *dto.UserCredential) (string, error) {
	user, err := u.GetUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *usecase) GetAllUsers() ([]model.User, error) {
	return u.repository.GetAllUsers()
}

func (u *usecase) UpdateUser(user *model.User, input *dto.NewUser) (*model.User, error) {
	password := []byte(input.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = string(hashedPassword)
	return u.repository.UpdateUser(user)
}

func (u *usecase) DeleteUser(user *model.User) (*model.User, error) {
	err := u.repository.DeleteUser(user)
	user.DeletedAt = &time.Time{}
	return user, err
}
