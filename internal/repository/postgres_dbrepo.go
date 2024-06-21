package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/models"
)

type postgresDBRepo struct {
	db *sql.DB
}

func New(cfg config.ConfigDB) DatabaseRepo {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := connectToDB(connStr)
	if err != nil {
		log.Fatal("Couldn't connect to DB: %v", err)
	}

	return &postgresDBRepo{
		db: db,
	}
}

func connectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("успешное подключение к БД!")
	return db, nil
}

const dbTimeout = time.Second * 3

func (m *postgresDBRepo) Connection() *sql.DB {
	return m.db //поправить на маленькие
}

func (m *postgresDBRepo) AllBooks() ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select b.id, name_book, concat_ws(' ', name_author, sirname_author) as author, year_book, isbn
        from books b
        inner join authors a on b.author_id = a.id
        order by name_book
    `

	rows, err := m.db.QueryContext(ctx, query)
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

func (m *postgresDBRepo) OneBook(id int) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select b.id, name_book, concat_ws(' ', name_author, sirname_author) as author, year_book, isbn
        from books b
        inner join authors a on b.author_id = a.id
        where b.id =$1
    `
	row := m.db.QueryRowContext(ctx, query, id)
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

func (m *postgresDBRepo) InsertBook(book models.BookID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        insert into books(name_book, author_id, year_book,isbn)
		values($1, $2, $3, $4) returning id
    `

	var newID int

	log.Println(book)

	err := m.db.QueryRowContext(ctx, query,
		book.Title,
		book.AuthorID,
		book.Year,
		book.ISBN,
	).Scan(&newID)

	if err != nil {
		log.Println("Ошибка тут!!!!!")
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) UpdateBook(book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        update books set name_book = $1, author_id = $2, year_book = $3, isbn =$4
		where id = $5
    `

	_, err := m.db.ExecContext(ctx, query,
		book.Title,
		book.Author,
		book.Year,
		book.ISBN,
		book.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) DeleteBook(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        delete from books where id = $1
    `

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

//здесь и далее действия с авторами

func (m *postgresDBRepo) AllAuthors() ([]*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        select id, name_author, sirname_author, biography, birthday
		from authors
        order by name_author
    `

	rows, err := m.db.QueryContext(ctx, query)
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

func (m *postgresDBRepo) OneAuthor(id int) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select id, name_author, sirname_author, biography, birthday
		from authors
        where id =$1
    `
	row := m.db.QueryRowContext(ctx, query, id)
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

func (m *postgresDBRepo) InsertAuthor(author models.Author) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        insert into authors(name_author, sirname_author, biography, birthday)
		values($1, $2, $3, $4) returning id
    `

	t, err := time.Parse("2006-01-02", author.Birthday)
	if err != nil {
		return 0, err
	}

	var newID int

	err = m.db.QueryRowContext(ctx, query,
		author.NameAuthor,
		author.SirnameAuthor,
		author.Biography,
		t,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) UpdateAuthor(author models.Author) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        update authors set name_author = $1, sirname_author = $2, biography = $3, birthday = $4
		where id = $5
    `

	fmt.Println(author.ID, author.Birthday)

	_, err := m.db.ExecContext(ctx, query,
		author.NameAuthor,
		author.SirnameAuthor,
		author.Biography,
		author.Birthday,
		author.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) DeleteAuthor(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        delete from authors where id = $1
    `

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) UpdateAuthorAndBook(author models.Author, book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	query := `
        update authors set name_author = $1, sirname_author = $2, biography = $3, birthday = $4
		where id = $5
    `

	_, err = m.db.ExecContext(ctx, query,
		author.NameAuthor,
		author.SirnameAuthor,
		author.Biography,
		author.Birthday,
		author.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	query = `
        update books set name_book = $1, author_id = $2, year_book = $3, isbn =$4
		where id = $5
    `

	_, err = m.db.ExecContext(ctx, query,
		book.Title,
		book.Author,
		book.Year,
		book.ISBN,
		book.ID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
