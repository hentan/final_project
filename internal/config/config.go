package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envVarDbHost = "DB_HOST"
const envVarDbPort = "DB_PORT"
const envVarDbUser = "DB_USER"
const envVarDbPassword = "DB_Password"
const envVarDbName = "DB_NAME"

// сделать одну структуру на 2 поля 1 - db, 2- server
type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Server struct {
	Port string
}

func NewConfigDB(path string) ConfigDB {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file")
	}

	return ConfigDB{
		Host:     getEnv(envVarDbHost, "db"),
		Port:     getEnv(envVarDbPort, "5432"),
		User:     getEnv(envVarDbUser, "postgres"),
		Password: getEnv(envVarDbPassword, "postgres"),
		Name:     getEnv(envVarDbName, "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewServer() Server {
	port := getEnv("APP_PORT", "8080")
	port = fmt.Sprintf(":%s", port)
	return Server{
		Port: port,
	}
}
