package api

import (
	"encoding/json"
	"go-ticket-tracker/internal/database"
	"net/http"
	"strconv"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// fallback if encoding fails, though rare with simple structs
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// TicketsHandler handles /tickets (GET for list, POST for create)
func TicketsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTickets(w, r)
	case http.MethodPost:
		createTicket(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// TicketHandler handles /tickets/ (GET for get, PUT for update)
func TicketHandler(w http.ResponseWriter, r *http.Request) {
	// Expect path /tickets/{id}
	// Verify ID exists
	idStr := strings.TrimPrefix(r.URL.Path, "/tickets/")
	if idStr == "" || idStr == "/" {
		writeError(w, http.StatusBadRequest, "Missing ticket ID")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ticket ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTicket(w, r, id)
	case http.MethodPut:
		updateTicket(w, r, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func listTickets(w http.ResponseWriter, r *http.Request) {
	excludeArea := r.URL.Query().Get("exclude_area")
	excludeStatus := r.URL.Query().Get("exclude_status")

	tickets, err := database.ListTickets(excludeArea, excludeStatus)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list tickets")
		return
	}
	// Ensure empty list is [] not null
	if tickets == nil {
		tickets = []database.Ticket{}
	}
	writeJSON(w, http.StatusOK, tickets)
}

func createTicket(w http.ResponseWriter, r *http.Request) {
	var ticket database.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := database.CreateTicket(&ticket); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create ticket")
		return
	}

	writeJSON(w, http.StatusCreated, ticket)
}

func getTicket(w http.ResponseWriter, r *http.Request, id int) {
	ticket, err := database.GetTicket(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "Ticket not found")
		return
	}
	writeJSON(w, http.StatusOK, ticket)
}

func updateTicket(w http.ResponseWriter, r *http.Request, id int) {
	var ticket database.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := database.UpdateTicket(id, &ticket); err != nil {
		writeError(w, http.StatusNotFound, "Ticket not found or update failed")
		return
	}

	// Fetch fresh to return
	updated, err := database.GetTicket(id)
	if err == nil {
		writeJSON(w, http.StatusOK, updated)
	} else {
		// Fallback to what we have if fetch fails
		writeJSON(w, http.StatusOK, ticket)
	}
}
