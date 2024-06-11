package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

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
	var app application

	envFilePath := "configs/api.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// create connection string from env
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	flag.StringVar(&app.DSN, "dsn", connStr, "Postgres connection string")

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
