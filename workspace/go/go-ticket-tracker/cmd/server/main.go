package main

import (
	"fmt"
	"go-ticket-tracker/internal/api"
	"go-ticket-tracker/internal/database"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize Database
	dbPath := "tickets.db"
	if err := database.InitDB(dbPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Printf("Database initialized at %s", dbPath)

	// Setup Router
	mux := http.NewServeMux()
	
	// Exact match for /tickets
	mux.HandleFunc("/tickets", api.TicketsHandler)
	// Prefix match for /tickets/ (handles IDs)
	mux.HandleFunc("/tickets/", api.TicketHandler)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
