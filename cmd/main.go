package main

import (
	"log"

	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/repository"
)

func main() {
	//create connection string and parse it
	envFilePath := "configs/api.env"
	dbCfg := config.NewConfigDB(envFilePath)
	repo := repository.New(dbCfg)
	appCfg := config.NewServer()
	app := handlers.New(repo, appCfg) //заполнить поле, чтобы передавалось
	// start database

	err := app.Start(handlers.Routes(app))
	if err != nil {
		log.Fatal("start application failed %v", err)
	}
}
