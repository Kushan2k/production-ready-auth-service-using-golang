package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
	SERVER_PORT string
	JWT_SECRET  string
	DEBUG       bool
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
		return nil, err
	}

	return &Config{
		DB_Host:     getEnv("DB_HOST", "localhost"),
		DB_Port:     getEnv("DB_PORT", "3306"),
		DB_User:     getEnv("DB_USER", "root"),
		DB_Password: getEnv("DB_PASSWORD", ""),
		DB_Name:     getEnv("DB_NAME", "auth_api"),
		SERVER_PORT: getEnv("SERVER_PORT", "8080"),
		JWT_SECRET:  getEnv("JWT_SECRET", "secret"),
		DEBUG:       getEnv("DEBUG", "true") == "true",
	}, nil
}


func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
