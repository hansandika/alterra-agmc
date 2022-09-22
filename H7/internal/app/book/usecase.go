package book

import (
	"errors"
	"net/http"

	"github.com/hansandika/internal/dto"
	"github.com/hansandika/internal/factory"
	"github.com/hansandika/internal/model"
	"github.com/hansandika/internal/repository"
	"github.com/hansandika/pkg/constant"
	"github.com/hansandika/pkg/util/response"
)

type UsecaseInterface interface {
	CreateNewBook(input *dto.NewBook) (*dto.BookResponse, *response.ErrorResponse)
	GetBookById(id int) (*dto.BookResponse, *response.ErrorResponse)
	GetAllBooks() ([]*dto.BookResponse, *response.ErrorResponse)
	UpdateBook(id int, input *dto.NewBook) (*dto.BookResponse, *response.ErrorResponse)
	DeleteBook(id int) (*dto.BookResponse, *response.ErrorResponse)
}

type usecase struct {
	BookRepository repository.BookRepositoryInterface
}

func NewUsecase(f *factory.Factory) UsecaseInterface {
	return &usecase{
		BookRepository: f.BookRepository,
	}
}

func (u *usecase) CreateNewBook(input *dto.NewBook) (*dto.BookResponse, *response.ErrorResponse) {
	var result *dto.BookResponse

	book, err := u.BookRepository.CreateNewBook(&model.Book{
		Title:         input.Title,
		Description:   input.Description,
		Author:        input.Author,
		YearPublished: input.YearPublished,
	})
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.BookResponse{
		ID:            int(book.ID),
		Title:         book.Title,
		Description:   book.Description,
		Author:        book.Author,
		YearPublished: book.YearPublished,
	}

	return result, nil
}

func (u *usecase) GetBookById(id int) (*dto.BookResponse, *response.ErrorResponse) {
	var result *dto.BookResponse

	book, err := u.BookRepository.GetBookById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("Book not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.BookResponse{
		ID:            int(book.ID),
		Title:         book.Title,
		Description:   book.Description,
		Author:        book.Author,
		YearPublished: book.YearPublished,
	}

	return result, nil
}

func (u *usecase) GetAllBooks() ([]*dto.BookResponse, *response.ErrorResponse) {
	var result []*dto.BookResponse

	books, err := u.BookRepository.GetAllBooks()
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	for _, book := range books {
		result = append(result, &dto.BookResponse{
			ID:            int(book.ID),
			Title:         book.Title,
			Description:   book.Description,
			Author:        book.Author,
			YearPublished: book.YearPublished,
		})
	}

	return result, nil
}

func (u *usecase) UpdateBook(id int, input *dto.NewBook) (*dto.BookResponse, *response.ErrorResponse) {
	var result *dto.BookResponse

	book, err := u.BookRepository.GetBookById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("Book not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	if input.Title != "" {
		book.Title = input.Title
	}

	if input.Description != "" {
		book.Description = input.Description
	}

	if input.Author != "" {
		book.Author = input.Author
	}

	if input.YearPublished != 0 {
		book.YearPublished = input.YearPublished
	}

	book, err = u.BookRepository.UpdateBook(book)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.BookResponse{
		ID:            int(book.ID),
		Title:         book.Title,
		Description:   book.Description,
		Author:        book.Author,
		YearPublished: book.YearPublished,
	}

	return result, nil
}

func (u *usecase) DeleteBook(id int) (*dto.BookResponse, *response.ErrorResponse) {
	var result *dto.BookResponse

	book, err := u.BookRepository.GetBookById(id)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, response.NewErrorResponse(http.StatusNotFound, errors.New("Book not found"))
		}
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	err = u.BookRepository.DeleteBook(book)
	if err != nil {
		return result, response.NewErrorResponse(http.StatusInternalServerError, err)
	}

	result = &dto.BookResponse{
		ID:            int(book.ID),
		Title:         book.Title,
		Description:   book.Description,
		Author:        book.Author,
		YearPublished: book.YearPublished,
	}

	return result, nil
}
