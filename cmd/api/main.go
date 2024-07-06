package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/config"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/routes"
)

func main() {
	config.LoadConfig("../../.env")

	db.ConnectDB()

	router := routes.SetupRouter()

	log.Printf("Starting server on port %d...", config.AppConfig.AppPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.AppPort), router)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
