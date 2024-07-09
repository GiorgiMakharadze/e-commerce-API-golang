package main

import (
	"log"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/config"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/models"
)

func init() {
	config.LoadConfig("../../.env")
	db.ConnectDB()
}

func main() {
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	} else {
		log.Println("Migration completed successfully")
	}
}
