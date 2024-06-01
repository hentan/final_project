package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hentan/final_project/internal/repository"
	"github.com/hentan/final_project/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	//конфиг
	var app application
	//флаги
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=final_project sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	//база данных
	conn, err := app.connectToDB()
	if err != nil {
		log.Println("не удалось подключить к БД!")
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	app.Domain = "example.com"

	log.Println("Старт приложения на порту:", port)

	//старт сервера
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
