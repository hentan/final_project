package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	//конфиг
	var app application
	//флаги

	//база данных
	app.Domain = "example.com"

	log.Println("Старт приложения на порту:", port)

	//старт сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
