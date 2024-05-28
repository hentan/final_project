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

func (m *PostgresDBRepo) OneBook(id int) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select b.id, name_book, concat_ws(' ', name_author, sirname_author) as author, year_book, isbn
        from books b
        inner join authors a on b.author_id = a.id
        where b.id =$1
    `
	row := m.DB.QueryRowContext(ctx, query, id)
	var book models.Book
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ISBN,
	)
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (m *PostgresDBRepo) AllAuthors() ([]*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select id, name_author, sirname_author, biography, birthday
		from authors
        order by name_author
    `

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*models.Author

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.ID,
			&author.NameAuthor,
			&author.SirnameAuthor,
			&author.Biography,
			&author.Birthday,
		)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return authors, nil
}

func (m *PostgresDBRepo) OneAuthor(id int) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select id, name_author, sirname_author, biography, birthday
		from authors
        where id =$1
    `
	row := m.DB.QueryRowContext(ctx, query, id)
	var author models.Author
	err := row.Scan(
		&author.ID,
		&author.NameAuthor,
		&author.SirnameAuthor,
		&author.Biography,
		&author.Birthday,
	)
	if err != nil {
		return nil, err
	}
	return &author, err
}
