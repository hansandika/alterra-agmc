package book

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
	f              = factory.Factory{BookRepository: repository.InitBookRepository(db)}
	controllerTest = NewController(&f)
)

func TestControllerCreateNewBookInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/api/v1/books")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.CreateNewBook(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Request body can't be empty")
	}
}

func TestControllerCreateNewBookSuccess(t *testing.T) {
	newBook := &dto.NewBook{
		Title:         "The Lord of the Rings",
		Description:   "This is book about the rings",
		Author:        "J. R. R. Tolkien",
		YearPublished: 1954,
	}
	payload, err := json.Marshal(newBook)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/books")
	c.Request().Header.Set("Content-Type", "application/json")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.CreateNewBook(c)) {
		asserts.Equal(201, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "title")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
		asserts.Contains(body, "description")
	}
}

func TestControllerGetBookByIdSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.GetBookById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "title")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
		asserts.Contains(body, "description")
	}
}

func TestControllerGetBookByIdNotFound(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.GetBookById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Book not found")
	}
}

func TestControllerGetAllBookSucess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/books")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(controllerTest.GetAllBooks(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestControllerUpdateBookNotFound(t *testing.T) {
	newBook := &dto.NewBook{
		Title:         "The Lord of the Rings",
		Description:   "This is book about the rings",
		Author:        "J. R. R. Tolkien",
		YearPublished: 1954,
	}
	payload, err := json.Marshal(newBook)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/books/:id")
	c.Request().Header.Set("Content-Type", "application/json")

	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.UpdateBookById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		fmt.Println(body)
		asserts.Contains(body, "Book not found")
	}
}

func TestControllerUpdateBookSuccess(t *testing.T) {
	newBook := &dto.NewBook{
		Title:         "The Lord of the Rings Update",
		Description:   "This is book about the rings Update",
		Author:        "J. R. R. Tolkien Update",
		YearPublished: 2000,
	}
	payload, err := json.Marshal(newBook)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	c.SetPath("/api/v1/books/:id")
	c.Request().Header.Set("Content-Type", "application/json")

	c.SetParamNames("id")
	c.SetParamValues("207")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.UpdateBookById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}

func TestControllerDeleteBookNotFound(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.DeleteBookById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Book not found")
	}
}

func TestControllerDeleteBookSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetPath("/api/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("213")

	asserts := assert.New(t)
	// testing
	if asserts.NoError(controllerTest.DeleteBookById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "title")
		asserts.Contains(body, "description")
		asserts.Contains(body, "author")
		asserts.Contains(body, "year_published")
	}
}
