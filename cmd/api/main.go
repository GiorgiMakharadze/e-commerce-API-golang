package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiorgiMakharadze/e-commerce-API-golang/config"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/db"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/internal/auth/handler"
	"github.com/GiorgiMakharadze/e-commerce-API-golang/routes"
	"github.com/gorilla/sessions"
)

func main() {
	config.LoadConfig("../../.env")

	db.ConnectDB()

	router := routes.SetupRouter()

	handler.Store = sessions.NewCookieStore([]byte(config.AppConfig.Session_key))
	handler.Store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
	}

	log.Printf("Starting server on port %d...", config.AppConfig.AppPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.AppPort), router)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
