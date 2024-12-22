package config

import (
	"fmt"
	"os"

	"github.com/hentan/final_project/internal/logger"
	"github.com/joho/godotenv"
)

const (
	envVarDbHost      = "DB_HOST"
	envVarDbPort      = "DB_PORT"
	envVarDbUser      = "DB_USER"
	envVarDbPassword  = "DB_Password"
	envVarDbName      = "DB_NAME"
	envVarKafkaBroker = "KAFKA_BROKER"
	envVarRedisAddres = "REDIS_ADDRESS"
)

type Kafka struct {
	Brokers []string
	Topic   string
}

type Redis struct {
	Addr     string
	Password string
}

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
	Kafka      Kafka
	Redis      Redis
}

func NewConfig(path string) Config {
	newLogger := logger.GetLogger()
	err := godotenv.Load(path)
	if err != nil {
		newLogger.Error("Error loading .env file")
	}

	port := getEnv("APP_PORT", "8080")
	port = fmt.Sprintf(":%s", port)
	kafkaBroker := getEnv(envVarKafkaBroker, "broker:29092")

	return Config{
		DBHost:     getEnv(envVarDbHost, "db"),
		DBPort:     getEnv(envVarDbPort, "5432"),
		DBUser:     getEnv(envVarDbUser, "postgres"),
		DBPassword: getEnv(envVarDbPassword, "postgres"),
		DBName:     getEnv(envVarDbName, "postgres"),
		AppPort:    port,
		Kafka: Kafka{
			Brokers: []string{kafkaBroker},
			Topic:   "errors_from_handlers",
		},
		Redis: Redis{
			Addr:     getEnv(envVarRedisAddres, "localhost:6379"),
			Password: "",
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
