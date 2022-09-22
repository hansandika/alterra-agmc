package user

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestControllerUserGetAllUsersSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/users")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.GetAllUsers(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
	}
}

func TestControllerUserGetUserByIdSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "2")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.GetUserById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
	}
}

func TestControllerUserGetUserByIdUnAuthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "3")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.GetUserById(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "This action is unauthorized")
	}
}

func TestControllerUserUpdateUserSuccess(t *testing.T) {
	newUser := &dto.NewUser{
		Name:     "Hans",
		Email:    "william@gmail.com",
		Password: "william02",
	}
	payload, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "2")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(controllerTest.UpdateUserById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
		asserts.Contains(body, "Update user success")
	}
}

func TestControllerUserUpdateUserUnauthorized(t *testing.T) {
	newUser := &dto.NewUser{
		Name:     "Hans",
		Email:    "william@gmail.com",
		Password: "william02",
	}
	payload, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "4")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(controllerTest.UpdateUserById(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "This action is unauthorized")
	}
}

func TestControllerUserDeleteUserSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("8")

	c.Request().Header.Add("X-Header-UserId", "8")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(controllerTest.DeleteUserById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Delete user success")
	}
}

func TestControllerUserDeleteUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("8")

	c.Request().Header.Add("X-Header-UserId", "9")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(controllerTest.DeleteUserById(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "This action is unauthorized")
	}
}
