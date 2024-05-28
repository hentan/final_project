package repository

import (
	"database/sql"

	"github.com/hentan/final_project/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllBooks() ([]*models.Book, error)
	OneBook(id int) (*models.Book, error)
	AllAuthors() ([]*models.Author, error)
	OneAuthor(id int) (*models.Author, error)
}
