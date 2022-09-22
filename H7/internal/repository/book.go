package repository

import (
	"github.com/hansandika/internal/model"
	"github.com/jinzhu/gorm"
)

type BookRepositoryInterface interface {
	CreateNewBook(book *model.Book) (*model.Book, error)
	GetBookById(id int) (*model.Book, error)
	GetAllBooks() ([]model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	DeleteBook(book *model.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func InitBookRepository(db *gorm.DB) BookRepositoryInterface {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) CreateNewBook(book *model.Book) (*model.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return nil, err
	}
	book, err = r.GetBookById(int(book.ID))
	return book, err
}

func (r *bookRepository) GetBookById(id int) (*model.Book, error) {
	var book model.Book
	err := r.db.Find(&book, id).Error
	return &book, err
}

func (r *bookRepository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) UpdateBook(book *model.Book) (*model.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *bookRepository) DeleteBook(book *model.Book) error {
	err := r.db.Delete(&book).Error
	return err
}
