package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
	AppPort     int
	Secret      string
	Session_key string
	Csrf_key    string
}

var AppConfig *Config

func LoadConfig(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s: %v", filePath, err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("Invalid APP_PORT: %v", err)
	}

	AppConfig = &Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      dbPort,
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		AppPort:     appPort,
		Secret:      os.Getenv("SECRET"),
		Session_key: os.Getenv("SESSION_KEY"),
		Csrf_key:    os.Getenv("CSRF_KEY"),
	}
}
