package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/hansandika/database"
	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/internal/mocks"
	"github.com/hansandika/internal/repository"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db             = database.GetConnection()
	echoMock       = mocks.EchoMock{E: echo.New()}
	f              = factory.Factory{UserRepository: repository.InitUserRepository(db)}
	controllerTest = NewController(&f)
)

func TestAuthLoginByEmailAndPasswordInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/api/v1/auth/login")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Request body can't be empty")
	}
}

func TestAuthLoginByEmailAndPasswordSuccess(t *testing.T) {
	userCred := &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "william02",
	}
	payload, err := json.Marshal(userCred)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/auth/login")
	c.Request().Header.Set("Content-Type", "application/json")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.LoginByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
		asserts.Contains(body, "token")
	}
}

func TestAuthLoginByEmailAndPasswordFailed(t *testing.T) {
	userCred := &dto.UserCredential{
		Email:    "william@gmail.com",
		Password: "william",
	}
	payload, err := json.Marshal(userCred)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/auth/login")
	c.Request().Header.Set("Content-Type", "application/json")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid email or password")
	}
}

func TestAuthRegisterByEmailAndPasswordUserAlreadyExists(t *testing.T) {
	newUser := &dto.NewUser{
		Name:     "William",
		Email:    "william@gmail.com",
		Password: "ayamgoreng",
	}
	payload, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/auth/register")
	c.Request().Header.Set("Content-Type", "application/json")
	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.RegisterUserByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Email already exists")
	}
}

func TestAuthRegisterByEmailAndPasswordSuccess(t *testing.T) {
	newUser := &dto.NewUser{
		Name:     "michael",
		Email:    "michael@gmail.com",
		Password: "ayamgoreng",
	}

	payload, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/auth/register")
	c.Request().Header.Set("Content-Type", "application/json")
	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.RegisterUserByEmailAndPassword(c)) {
		asserts.Equal(201, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "Register success")
	}
}
