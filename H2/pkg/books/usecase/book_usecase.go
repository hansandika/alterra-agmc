package usecase

import (
	"github.com/hansandika/pkg/books/dto"
	"github.com/hansandika/pkg/books/model"
	"github.com/hansandika/pkg/books/repository"
)

type UsecaseInterface interface {
	CreateNewBook(input *dto.NewBook) (*model.Book, error)
	GetBookById(id int) (*model.Book, error)
	GetAllBooks() ([]model.Book, error)
	UpdateBook(id int, input *dto.NewBook) (*model.Book, error)
	DeleteBook(id int) (*model.Book, error)
}

type usecase struct {
	repository repository.RepositoryInterface
}

func InitUsecase(repository repository.RepositoryInterface) UsecaseInterface {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateNewBook(input *dto.NewBook) (*model.Book, error) {
	return u.repository.CreateNewBook(&model.Book{
		Title:         input.Title,
		Description:   input.Description,
		Author:        input.Author,
		YearPublished: input.YearPublished,
	})
}

func (u *usecase) GetBookById(id int) (*model.Book, error) {
	return u.repository.GetBookById(id)
}

func (u *usecase) GetAllBooks() ([]model.Book, error) {
	return u.repository.GetAllBooks()
}

func (u *usecase) UpdateBook(id int, input *dto.NewBook) (*model.Book, error) {
	book, err := u.repository.GetBookById(id)
	if err != nil {
		return nil, err
	}
	book.Title = input.Title
	book.Description = input.Description
	book.Author = input.Author
	book.YearPublished = input.YearPublished
	return u.repository.UpdateBook(book)
}

func (u *usecase) DeleteBook(id int) (*model.Book, error) {
	book, err := u.repository.GetBookById(id)
	if err != nil {
		return nil, err
	}
	err = u.repository.DeleteBook(book)
	return book, err
}
