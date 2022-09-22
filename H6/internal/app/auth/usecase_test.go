package auth

import (
	"testing"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/stretchr/testify/assert"
)

var (
	usecaseTest = NewUsecase(factory.NewFactory())
)

func TestAuthUsecaseLoginByEmailAndPasswordSuccess(t *testing.T) {
	asserts := assert.New(t)

	payload := &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "william02",
	}

	res, err := usecaseTest.LoginByEmailAndPassword(payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
}

func TestAuthUsecaseLoginByEmailAndPasswordNotFound(t *testing.T) {
	asserts := assert.New(t)

	payload := &dto.UserCredential{
		Email:    "mmiawmiaw@gmail.com",
		Password: "asd",
	}

	_, err := usecaseTest.LoginByEmailAndPassword(payload)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "User not found")
	}
}

func TestAuthUsecaseLoginByEmailAndPasswordInvalidCredential(t *testing.T) {
	asserts := assert.New(t)

	payload := &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "william123",
	}

	_, err := usecaseTest.LoginByEmailAndPassword(payload)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "Invalid email or password")
	}
}

func TestAuthUsecaseRegisterUserByEmailAndPasswordSuccess(t *testing.T) {
	asserts := assert.New(t)

	payload := &dto.NewUser{
		Name:     "ciacia",
		Email:    "cia@gmail.com",
		Password: "cia02",
	}

	res, err := usecaseTest.RegisterUserByEmailAndPassword(payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
	asserts.Equal(payload.Name, res.Name)
}

func TestAuthUsecaseRegisterUserByEmailAndPasswordUserExists(t *testing.T) {
	asserts := assert.New(t)

	payload := &dto.NewUser{
		Name:     "ciacia",
		Email:    "cia@gmail.com",
		Password: "cia02",
	}
	_, err := usecaseTest.RegisterUserByEmailAndPassword(payload)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "Email already exists")
	}
}
