package factory

import (
	"github.com/hansandika/database"
	"github.com/hansandika/internal/repository"
)

type Factory struct {
	UserRepository repository.UserRepositoryInterface
	BookRepository repository.BookRepositoryInterface
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepository: repository.InitUserRepository(db),
		BookRepository: repository.InitBookRepository(db),
	}
}
