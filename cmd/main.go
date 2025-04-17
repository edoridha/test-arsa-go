package main

import (
	"log"
    "net/http"           // Untuk http.ListenAndServe
    "test-ara/db"        // Untuk db.InitDB()
    "test-ara/routes"    // Untuk routes.SetupRoutes()
    "github.com/go-chi/chi/v5" // Untuk chi.NewRouter()
	"github.com/go-chi/cors"
)

func main() {
    log.Println("Initializing database...")
    db.InitDB()

	log.Println("Seeding database...")
	db.Seed()

    log.Println("Setting up routes...")
    r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"}, 
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, 
    }))

    routes.SetupRoutes(r)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}