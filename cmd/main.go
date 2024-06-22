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
	cfg := config.NewConfigDB(envFilePath)
	repo := repository.New(cfg)
	app := handlers.New(repo, cfg)
	// start database
	err := app.Start(handlers.Routes(app))
	if err != nil {
		log.Fatal("start application failed %v", err)
	}
}
