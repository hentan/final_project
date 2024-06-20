package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func (c *ConfigDB) NewConfigDB(path string) string {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file")
	}

	conn := ConfigDB{
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "postgres"),
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conn.DBUser, conn.DBPassword, conn.DBHost, conn.DBPort, conn.DBName)

	return connStr
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
