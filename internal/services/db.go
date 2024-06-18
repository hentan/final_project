package services

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB(dsn string) (*sql.DB, error) {
	connection, err := OpenDB(dsn)
	if err != nil {
		return nil, err
	}
	log.Println("успешное подключение к БД!")
	return connection, nil
}
