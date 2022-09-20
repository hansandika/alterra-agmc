package usecase_test

import (
	"fmt"
	"testing"

	"github.com/hansandika/config"
	"github.com/hansandika/mocks"
	"github.com/hansandika/pkg/books/dto"
	"github.com/hansandika/pkg/books/model"
	"github.com/hansandika/pkg/books/repository"
	"github.com/hansandika/pkg/books/usecase"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db        = config.InitDB()
	echoMock  = mocks.EchoMock{E: echo.New()}
	repo      = repository.InitRepository(db)
	uc        = usecase.InitUsecase(repo)
	bookInput = dto.NewBook{
		Title:         "Test Update",
		Description:   "Test Update",
		Author:        "Hans",
		YearPublished: 2020,
	}
)

func TestGetAllBookSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	for _, val := range res {
		asserts.NotEmpty(val.ID)
		asserts.NotEmpty(val.Title)
		asserts.NotEmpty(val.Description)
		asserts.NotEmpty(val.Author)
		asserts.NotEmpty(val.YearPublished)
	}
}

func TestGetBookByIdSuccess(t *testing.T) {
	res, err := uc.GetBookById(4)
	if err != nil {
		t.Fatal(err)
	}
	asserts := assert.New(t)
	asserts.Equal(uint(4), res.ID)
}

func TestGetBookNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.GetBookById(100)
	if err != nil {
		asserts.Equal(err.Error(), "record not found")
	}
}

func TestCreateBookSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.CreateNewBook(&bookInput)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(bookInput.Title, res.Title)
	asserts.Equal(bookInput.Description, res.Description)
	asserts.Equal(bookInput.Author, res.Author)
	asserts.Equal(bookInput.YearPublished, res.YearPublished)
}

func TestCreateBookInvalidPayload(t *testing.T) {
	asserts := assert.New(t)
	bookInput.Title = ""
	_, err := uc.CreateNewBook(&bookInput)
	if err != nil {
		asserts.Equal(err.Error(), "Title can't be empty")
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.UpdateBook(&model.Book{
		Model:       gorm.Model{ID: 11},
		Title:       "Test",
		Description: "Test",
	}, &bookInput)

	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(bookInput.Title, res.Title)
	asserts.Equal(bookInput.Description, res.Description)
	asserts.Equal(bookInput.Author, res.Author)
	asserts.Equal(bookInput.YearPublished, res.YearPublished)
}

func TestUpdateBookNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.UpdateBook(&model.Book{
		Model:       gorm.Model{ID: 200},
		Title:       "Test",
		Description: "Test",
	}, &bookInput)

	if err != nil {
		fmt.Println(err.Error())
		asserts.Equal(err.Error(), "record not found")
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	asserts := assert.New(t)
	res, err := uc.DeleteBook(&model.Book{
		Model: gorm.Model{ID: 200},
	})

	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestDeleteBookNotFound(t *testing.T) {
	asserts := assert.New(t)
	_, err := uc.DeleteBook(&model.Book{
		Model: gorm.Model{ID: 1},
	})

	if err != nil {
		fmt.Println(err.Error())
		asserts.Equal(err.Error(), "record not found")
	}
}
