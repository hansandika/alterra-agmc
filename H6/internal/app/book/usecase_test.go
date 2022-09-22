package book

import (
	"testing"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/stretchr/testify/assert"
)

var (
	usecaseTest = NewUsecase(factory.NewFactory())
)

func TestBookUsecaseGetAllBooks(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}

	for _, val := range res {
		asserts.NotEmpty(val.ID)
	}
}

func TestBookUsecaseGetBookByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.GetBookById(4)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestBookUsecaseGetBookByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := usecaseTest.GetBookById(1)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "Book not found")
	}
}

func TestBookUsecaseCreateBookSuccess(t *testing.T) {
	asserts := assert.New(t)
	payload := &dto.NewBook{
		Title:         "The Lord of the Rings",
		Description:   "The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien.",
		Author:        "J. R. R. Tolkien",
		YearPublished: 1954,
	}
	res, err := usecaseTest.CreateNewBook(payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestBookUsecaseUpdateBookByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	payload := &dto.NewBook{
		Title:         "dummy",
		Description:   "dummy",
		Author:        "dummy",
		YearPublished: 2020,
	}
	_, err := usecaseTest.UpdateBook(1, payload)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "Book not found")
	}
}

func TestBookUsecaseUpdateBookByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	payload := &dto.NewBook{
		Title:         "dummy",
		Description:   "dummy",
		Author:        "dummy",
		YearPublished: 2020,
	}
	res, err := usecaseTest.UpdateBook(211, payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}

func TestBookUsecaseDeleteBookByIdNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := usecaseTest.DeleteBook(1)
	if asserts.Error(err.ErrorMessage) {
		asserts.Equal(err.ErrorMessage.Error(), "Book not found")
	}
}

func TestBookUsecaseDeleteBookByIdSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := usecaseTest.DeleteBook(211)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
}
