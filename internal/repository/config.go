package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	dBHost     string
	dBPort     string
	dBUser     string
	dBPassword string
	dBName     string
}

type PortApp struct {
	port string
}

func (c *ConfigDB) NewConfigDB(path string) string {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file")
	}

	conn := ConfigDB{
		dBHost:     getEnv("DB_HOST", "db"),
		dBPort:     getEnv("DB_PORT", "5432"),
		dBUser:     getEnv("DB_USER", "postgres"),
		dBPassword: getEnv("DB_PASSWORD", "postgres"),
		dBName:     getEnv("DB_NAME", "postgres"),
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conn.dBUser, conn.dBPassword, conn.dBHost, conn.dBPort, conn.dBName)

	return connStr
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (p *PortApp) GetPortApp() string {
	portApp := PortApp{
		port: getEnv("APP_PORT", "8080"),
	}
	port := fmt.Sprintf(":%s", portApp.port)
	log.Println(port)
	return port
}
