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
	flag.StringVar(&app.DSN, "dsn", "postgres://postgres:postgres@db:5432/final_project?sslmode=disable", "Postgres connection string")
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
		log.Fatal("привет", err)
	}
}
