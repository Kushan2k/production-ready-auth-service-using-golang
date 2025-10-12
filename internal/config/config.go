package config

import (
	"log"
	"os"
	"strconv"

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

	MAIL_HOST     string
	MAIL_PORT     int
	MAIL_USERNAME string
	MAIL_PASSWORD string
}

var Envs=loadEnvs()

func LoadConfig() (*Config, error) {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using environment variables or defaults")
		
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
		MAIL_HOST:     getEnv("MAIL_HOST", ""),
		MAIL_PORT:    func() int {
			port, err := strconv.Atoi(getEnv("MAIL_PORT", "465"))
			if err != nil {
				return 465
			}
			return port
		}(),
		MAIL_USERNAME: getEnv("MAIL_USERNAME", ""),
		MAIL_PASSWORD: getEnv("MAIL_PASSWORD", ""),
	}, nil
}

func loadEnvs() Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return *cfg
}


func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
