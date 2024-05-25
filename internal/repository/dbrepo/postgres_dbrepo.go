package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/hentan/final_project/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllBooks() ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select b.id, name_book, concat_ws(' ', name_author, sirname_author) as author, year_book, isbn
        from books b
        inner join authors a on b.author_id = a.id
        order by name_book
    `

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.ISBN,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
