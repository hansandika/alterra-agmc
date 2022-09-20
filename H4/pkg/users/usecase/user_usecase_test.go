package usecase_test

import (
	"fmt"
	"testing"

	"github.com/hansandika/config"
	"github.com/hansandika/mocks"
	"github.com/hansandika/pkg/users/dto"
	"github.com/hansandika/pkg/users/model"
	"github.com/hansandika/pkg/users/repository"
	"github.com/hansandika/pkg/users/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db        = config.InitDB()
	echoMock  = mocks.EchoMock{E: echo.New()}
	repo      = repository.InitRepository(db)
	uc        = usecase.InitUsecase(repo)
	userInput = &dto.NewUser{
		Name:     "coki anwar update",
		Email:    "cokiAnwar@gmail.com",
		Password: "asd",
	}
	userCredential = &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "william02",
	}
	invalidUserCredential = &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "asd",
	}
)

func TestGetAllUserSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.GetAllUsers()
	if err != nil {
		t.Fatal(err)
	}
	for _, val := range res {
		asserts.NotEmpty(val.ID)
		asserts.NotEmpty(val.Name)
		asserts.NotEmpty(val.Email)
		asserts.NotEmpty(val.Password)
	}
}

func TestGetUserByIdSuccess(t *testing.T) {
	res, err := uc.GetUserById(1)
	if err != nil {
		t.Fatal(err)
	}
	asserts := assert.New(t)
	asserts.Equal(uint(1), res.ID)
}

func TestGetUserNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.GetUserById(100)
	if err != nil {
		asserts.Equal(err.Error(), "record not found")
	}
}

func TestLoginUserSuccess(t *testing.T) {
	res, err := uc.GetUserCredential(userCredential)
	if err != nil {
		t.Fatal(err)
	}
	asserts := assert.New(t)
	fmt.Println(res)
	asserts.NotContains(res, "crypto/bcrypt: hashedPassword is not the hash of the given password")
}

func TestLoginUserInvalidCredential(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.GetUserCredential(invalidUserCredential)
	if err != nil {
		fmt.Println(err.Error())
		asserts.Equal(err.Error(), "crypto/bcrypt: hashedPassword is not the hash of the given password")
	}
}

func TestCreateUserSuccess(t *testing.T) {
	res, err := uc.CreateNewUser(userInput)
	if err != nil {
		t.Fatal(err)
	}
	asserts := assert.New(t)
	asserts.NotEmpty(res.ID)
	asserts.NotEmpty(res.Name)
	asserts.NotEmpty(res.Email)
	asserts.NotEmpty(res.Password)
}

func TestCreateUserInvalidPayload(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.CreateNewUser(userInput)
	if err != nil {
		asserts.Equal(err.Error(), "pq: duplicate key value violates unique constraint \"users_email_key\"")
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	res, err := uc.UpdateUser(&model.User{
		Model: gorm.Model{ID: 1},
	}, userInput)
	if err != nil {
		t.Fatal(err)
	}
	asserts := assert.New(t)
	asserts.NotEmpty(res.ID)
	asserts.NotEmpty(res.Name)
	asserts.NotEmpty(res.Email)
	asserts.NotEmpty(res.Password)
}

func TestUpdateUserNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.UpdateUser(&model.User{
		Model: gorm.Model{ID: 100},
	}, userInput)
	if err != nil {
		asserts.Equal(err.Error(), "record not found")
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.DeleteUser(&model.User{
		Model: gorm.Model{ID: 11},
	})
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestDeleteUserNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.DeleteUser(&model.User{
		Model: gorm.Model{ID: 100},
	})
	if err != nil {
		fmt.Println(err.Error())
		asserts.Equal(err.Error(), "record not found")
	}
}
