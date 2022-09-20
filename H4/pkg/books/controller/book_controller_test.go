package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/hansandika/config"
	"github.com/hansandika/mocks"
	"github.com/hansandika/pkg/books/controller"
	"github.com/hansandika/pkg/books/repository"
	"github.com/hansandika/pkg/books/usecase"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db             = config.InitDB()
	echoMock       = mocks.EchoMock{E: echo.New()}
	repo           = repository.InitRepository(db)
	uc             = usecase.InitUsecase(repo)
	bookController = controller.InitControllerBook(uc)
)

func TestGetAllBookSucess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.GetAllBooks(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestCreateBookNoPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)

	c.SetPath("/api/v1/books")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.CreateNewBook(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Request body can't be empty")
	}
}

func TestCreateBookSuccess(t *testing.T) {
	jsonParam := `{"title":"book test","description":"test book description","author":"test author","year_published":2020}`
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBufferString(jsonParam))

	c.SetPath("/api/v1/books")
	c.Request().Header.Add("Content-Type", "application/json")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.CreateNewBook(c)) {
		asserts.Equal(201, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestGetBookByIdSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.GetBookById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestGetBookByIdNotFound(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.GetBookById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "record not found")
	}
}

func TestGetBookByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("abc")
	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.GetBookById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Invalid id param")
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	jsonParam := `{"title":"book test update","description":"test book description update","author":"test author update","year_published":2020}`
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBufferString(jsonParam))
	c.Request().Header.Add("Content-Type", "application/json")
	c.SetPath("/api/v1/books")
	c.SetParamNames("id")
	c.SetParamValues("5")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.UpdateBook(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestUpdateBookNotFound(t *testing.T) {
	jsonParam := `{"title":"book test update","description":"test book description update","author":"test author update","year_published":2020}`
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBufferString(jsonParam))
	c.Request().Header.Add("Content-Type", "application/json")
	c.SetPath("/api/v1/books")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.UpdateBook(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "record not found")
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("7")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.DeleteBook(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Book deleted succesfully")
	}
}

func TestDeleteBookNotFound(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(bookController.DeleteBook(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "record not found")
	}
}
