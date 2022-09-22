package user

import (
	"fmt"
	"testing"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/stretchr/testify/assert"
)

var (
	usecaseTest = NewUsecase(factory.NewFactory())
)

func TestUsecaseGetAllUserSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.GetAllUsers()
	if err != nil {
		t.Fatal(err)
	}

	for _, val := range res {
		asserts.NotEmpty(val.ID)
	}
}

func TestUsecaseGetUserByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := usecaseTest.GetUserById(500)
	if asserts.Error(err.ErrorMessage) {
		fmt.Println(err.ErrorMessage)
		asserts.Equal(err.ErrorMessage.Error(), "User not found")
	}
}

func TestUsecaseGetUserByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.GetUserById(4)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestUsecaseUpdateUserByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	user := &dto.NewUser{
		Name:     "william",
		Email:    "william@gmail.com",
		Password: "william02",
	}
	res, err := usecaseTest.UpdateUser(2, user)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestUsecaseUpdateUserByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	user := &dto.NewUser{
		Name:     "william",
		Email:    "william@gmail.com",
		Password: "william02",
	}
	_, err := usecaseTest.UpdateUser(404, user)
	if asserts.Error(err.ErrorMessage) {
		fmt.Println(err.ErrorMessage)
		asserts.Equal(err.ErrorMessage.Error(), "User not found")
	}
}

func TestUsecaseDeleteUserByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.DeleteUser(9)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestUsecaseDeleteUserByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := usecaseTest.DeleteUser(7)
	if asserts.Error(err.ErrorMessage) {
		fmt.Println(err.ErrorMessage)
		asserts.Equal(err.ErrorMessage.Error(), "User not found")
	}
}
