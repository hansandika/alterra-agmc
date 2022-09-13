package repository

import (
	"github.com/hansandika/pkg/books/model"
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	CreateNewBook(book *model.Book) (*model.Book, error)
	GetBookById(id int) (*model.Book, error)
	GetAllBooks() ([]model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	DeleteBook(book *model.Book) error
}

type repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateNewBook(book *model.Book) (*model.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	book, err = r.GetBookById(int(book.ID))
	return book, err
}

func (r *repository) GetBookById(id int) (*model.Book, error) {
	var book model.Book
	err := r.db.Find(&book, id).Error
	return &book, err
}

func (r *repository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) UpdateBook(book *model.Book) (*model.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) DeleteBook(book *model.Book) error {
	err := r.db.Delete(&book).Error
	return err
}
