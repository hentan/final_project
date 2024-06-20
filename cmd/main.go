package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/repository/dbrepo"
	"github.com/hentan/final_project/internal/services"
)

func main() {
	var app handlers.Application

	envFilePath := "configs/api.env"

	connStr := app.DBConf.NewConfigDB(envFilePath)
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	// create connection string from env

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
