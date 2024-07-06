package db

import (
	"fmt"
	"log"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.AppConfig.DBHost,
	config.AppConfig.DBPort,
	config.AppConfig.DBUser,
	config.AppConfig.DBPassword,
	config.AppConfig.DBName,

)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
}
