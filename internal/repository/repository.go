package repository

import (
	"database/sql"

	"github.com/hentan/final_project/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllBooks() ([]*models.Book, error)
	OneBook(id int) (*models.Book, error)
	InsertBook(book models.Book) (int, error)
	UpdateBook(book models.Book) error
	DeleteBook(id int) error
	AllAuthors() ([]*models.Author, error)
	OneAuthor(id int) (*models.Author, error)
	InsertAuthor(author models.Author) (int, error)
	UpdateAuthor(author models.Author) error
	DeleteAuthor(id int) error
}
