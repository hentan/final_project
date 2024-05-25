package repository

import (
	"database/sql"

	"github.com/hentan/final_project/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllBooks() ([]*models.Book, error)
}
