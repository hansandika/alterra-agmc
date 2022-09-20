package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/hansandika/config"
	"github.com/hansandika/mocks"
	"github.com/hansandika/pkg/users/controller"
	"github.com/hansandika/pkg/users/repository"
	"github.com/hansandika/pkg/users/usecase"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db             = config.InitDB()
	echoMock       = mocks.EchoMock{E: echo.New()}
	repo           = repository.InitRepository(db)
	uc             = usecase.InitUsecase(repo)
	userController = controller.InitControllerUser(uc)
	testBookID     = uint(1)
)

func TestGetAllUsersSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books")
	asserts := assert.New(t)
	if asserts.NoError(userController.GetAllUsers(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "ID")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
	}
}

func TestCreateNewUserInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)

	c.SetPath("/api/v1/books")
	asserts := assert.New(t)
	if asserts.NoError(userController.CreateNewUser(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Request body can't be empty")
	}
}

func TestCreateNewUserSuccess(t *testing.T) {
	jsonParam := `{
		"name": "Hans",
		"email": "hans@gmail.com",
		"password": "123456"
	}`

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBufferString(jsonParam))

	c.SetPath("/api/v1/users")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.CreateNewUser(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "ID")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
		asserts.Contains(body, "User added succesfully")
	}
}

func TestLoginUserSuccess(t *testing.T) {
	jsonParam := `{
		"email": "william@gmail.com",
		"password": "william02"
	}`
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBufferString(jsonParam))
	c.SetPath("/api/v1/users/login")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.LoginUser(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "token")
		asserts.Contains(body, "User logged in succesfully")
	}
}

func TestLoginUserInvalidPayload(t *testing.T) {
	jsonParam := `{
		"email": "william@gmail.com",
		"password": "aadasdasd"
	}`
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBufferString(jsonParam))
	c.SetPath("/api/v1/users/login")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.LoginUser(c)) {
		asserts.Equal(500, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
	}
}

func TestGetUserByIDSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "2")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.GetUserById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "User found")
	}
}

func TestGetUserByIDUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.GetUserById(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "this action is unauthorized")
	}
}

func TestUpdateUserUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.UpdateUser(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "this action is unauthorized")
	}
}

func TestUpdateUserInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "2")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.UpdateUser(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Request body can't be empty")
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	jsonParam := `{
		"email": "william@gmail.com",
		"password" : "william02",
		"name" : "william Huang"}`

	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBufferString(jsonParam))

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("2")

	c.Request().Header.Add("X-Header-UserId", "2")
	c.Request().Header.Add("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.UpdateUser(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "ID")
		asserts.Contains(body, "name")
		asserts.Contains(body, "email")
		asserts.Contains(body, "User updated succesfully")
	}
}

func TestDeleteUserUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("6")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.DeleteUser(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "this action is unauthorized")
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("6")

	c.Request().Header.Add("X-Header-UserId", "6")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.DeleteUser(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "User deleted succesfully")
	}
}

func TestDeleteeUserInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)

	c.SetPath("/api/v1/users/jwt")
	c.SetParamNames("id")
	c.SetParamValues("7")

	c.Request().Header.Add("X-Header-UserId", "7")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(userController.DeleteUser(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "record not found")
	}
}
