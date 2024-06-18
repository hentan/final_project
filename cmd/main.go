package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/repository/dbrepo"
	"github.com/hentan/final_project/internal/services"
)

func main() {
	var app handlers.Application

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
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	// create connection string from env
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	flag.StringVar(&app.DSN, "dsn", connStr, "Postgres connection string")

	flag.Parse()

	// start database
	conn, err := services.ConnectToDB(app.DSN)
	if err != nil {
		log.Println("не удалось подключить к БД!")
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	log.Println("Старт приложения на порту:", appPort)

	// start server
	err = http.ListenAndServe(fmt.Sprintf(":%d", appPort), handlers.Routes(&app))
	if err != nil {
		log.Fatal("привет", err)
	}
}
