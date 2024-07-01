package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	envVarDbHost     = "DB_HOST"
	envVarDbPort     = "DB_PORT"
	envVarDbUser     = "DB_USER"
	envVarDbPassword = "DB_Password"
	envVarDbName     = "DB_NAME"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func NewConfigDB(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := getEnv("APP_PORT", "8080")
	port = fmt.Sprintf(":%s", port)

	return Config{
		DBHost:     getEnv(envVarDbHost, "db"),
		DBPort:     getEnv(envVarDbPort, "5432"),
		DBUser:     getEnv(envVarDbUser, "postgres"),
		DBPassword: getEnv(envVarDbPassword, "postgres"),
		DBName:     getEnv(envVarDbName, "postgres"),
		AppPort:    port,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
