package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envVarDbName = "DB_NAME" //создать константы на все

// сделать одну структуру на 2 поля 1 - db, 2- server
// убрать db в названиях
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
		Host:     getEnv("DB_HOST", "db"), //вынести в константы
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Name:     getEnv("DB_NAME", "postgres"),
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
