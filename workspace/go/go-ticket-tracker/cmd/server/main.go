package main

import (
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

	// Wrap mux with CORS middleware
	corsMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(addr, corsMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
