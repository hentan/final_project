package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/repository/dbrepo"
	"github.com/hentan/final_project/internal/services"
)

func main() {
	var app handlers.Application

	envFilePath := "configs/api.env"
	connStr := app.DBConf.NewConfigDB(envFilePath)

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

	appPort := app.AppPort.GetPortApp()
	log.Println("Старт приложения на порту:", appPort)

	// start server
	err = http.ListenAndServe(appPort, handlers.Routes(&app))
	if err != nil {
		log.Fatal("connection error on 8080!", err)
	}
}
