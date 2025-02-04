package main

import (
	"log"

	_ "github.com/hentan/final_project/docs"
	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/handlers"
	"github.com/hentan/final_project/internal/kafka"
	"github.com/hentan/final_project/internal/logger"
	redispackage "github.com/hentan/final_project/internal/redis"
	"github.com/hentan/final_project/internal/repository"
)

//	@title			Simple Books API
//	@version		1.0
//	@description	This is a simple application for viewing authors and books
//  @host localhost:8080

func main() {
	//create connection string and parse it
	envFilePath := "configs/api.env"
	cfg := config.NewConfig(envFilePath)
	configLogger, _ := logger.NewConfigWithFormat("json")
	err := logger.InitGlobalLogger(configLogger)
	if err != nil {
		log.Fatal("Не удалось инициализировать глобальный логгер:", err)
	}
	newLogger := logger.GetLogger()
	redisClient := redispackage.NewRedisClient(cfg)
	kafkaProducer, err := kafka.NewKafkaProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	if err != nil {
		newLogger.Error("не удалось создать Kafka producer")
		return
	}

	repo := repository.New(cfg)
	if repo == nil {
		newLogger.Error("не удалось подключиться к базе данных!")
		return
	}
	app := handlers.New(repo, cfg, kafkaProducer, redisClient)

	// start database
	err = app.Start(handlers.Routes(app))
	if err != nil {
		newLogger.Error("не удалось запустить приложение")
		return
	}
}
