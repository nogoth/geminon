package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-ticket-tracker/internal/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupTestDB(t *testing.T) {
	err := database.InitDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to init in-memory DB: %v", err)
	}
}

func TestCreateTicket(t *testing.T) {
	setupTestDB(t)

	ticket := database.Ticket{
		Title:       "Test Ticket",
		Description: "A test description",
		Status:      "Open",
		AreaOfConcern: "Testing",
		AgentName:   "Tester",
		ResolutionMessage: "Pending",
	}
	body, _ := json.Marshal(ticket)
	req := httptest.NewRequest(http.MethodPost, "/tickets", bytes.NewReader(body))
	w := httptest.NewRecorder()

	TicketsHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status Created (201), got %v", resp.StatusCode)
	}

	var createdTicket database.Ticket
	json.NewDecoder(resp.Body).Decode(&createdTicket)

	if createdTicket.ID == 0 {
		t.Error("Expected ID to be set")
	}
	if createdTicket.Title != ticket.Title {
		t.Errorf("Expected Title %s, got %s", ticket.Title, createdTicket.Title)
	}
}

func TestListTicketsFilters(t *testing.T) {
	setupTestDB(t)

	// Create tickets
	tickets := []database.Ticket{
		{Title: "T1", Status: "Open", AreaOfConcern: "IT"},
		{Title: "T2", Status: "Closed", AreaOfConcern: "HR"},
		{Title: "T3", Status: "Open", AreaOfConcern: "HR"},
		{Title: "T4", Status: "Closed", AreaOfConcern: "IT"},
	}

	for _, ticket := range tickets {
		database.CreateTicket(&ticket)
	}

	// Test 1: Exclude Area "IT" -> Expect T2, T3 (HR)
	req := httptest.NewRequest(http.MethodGet, "/tickets?exclude_area=IT", nil)
	w := httptest.NewRecorder()
	TicketsHandler(w, req)
	
	var result []database.Ticket
	json.NewDecoder(w.Result().Body).Decode(&result)

	if len(result) != 2 {
		t.Errorf("Expected 2 tickets, got %d", len(result))
	}
	for _, tk := range result {
		if tk.AreaOfConcern == "IT" {
			t.Error("Found IT ticket despite exclusion")
		}
	}

	// Test 2: Exclude Status "Closed" -> Expect T1, T3 (Open)
	req = httptest.NewRequest(http.MethodGet, "/tickets?exclude_status=Closed", nil)
	w = httptest.NewRecorder()
	TicketsHandler(w, req)
	
	result = nil
	json.NewDecoder(w.Result().Body).Decode(&result)

	if len(result) != 2 {
		t.Errorf("Expected 2 tickets, got %d", len(result))
	}
	for _, tk := range result {
		if tk.Status == "Closed" {
			t.Error("Found Closed ticket despite exclusion")
		}
	}

	// Test 3: Exclude Both "IT" and "Closed" -> Expect T3 (Open & HR)
	// T1 (Open, IT) -> excluded by area
	// T2 (Closed, HR) -> excluded by status
	// T3 (Open, HR) -> kept
	// T4 (Closed, IT) -> excluded by both
	req = httptest.NewRequest(http.MethodGet, "/tickets?exclude_area=IT&exclude_status=Closed", nil)
	w = httptest.NewRecorder()
	TicketsHandler(w, req)
	
	result = nil
	json.NewDecoder(w.Result().Body).Decode(&result)

	if len(result) != 1 {
		t.Errorf("Expected 1 ticket, got %d", len(result))
	}
	if len(result) > 0 && result[0].Title != "T3" {
		t.Errorf("Expected T3, got %s", result[0].Title)
	}
}

func TestLongResolutionMessage(t *testing.T) {
	setupTestDB(t)

	longMsg := strings.Repeat("Lorem Ipsum ", 500) // ~6KB
	ticket := database.Ticket{
		Title:             "Long Message Ticket",
		ResolutionMessage: longMsg,
	}

	database.CreateTicket(&ticket)

	// Fetch via GET /tickets/{id}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/tickets/%d", ticket.ID), nil)
	w := httptest.NewRecorder()
	
	// We call TicketHandler directly but need to handle path logic?
	// TicketHandler uses strings.TrimPrefix(r.URL.Path, "/tickets/")
	// So path must be /tickets/{id}
	
	TicketHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected OK, got %v", resp.StatusCode)
	}

	var fetchedTicket database.Ticket
	json.NewDecoder(resp.Body).Decode(&fetchedTicket)

	if fetchedTicket.ResolutionMessage != longMsg {
		t.Error("Resolution message mismatch or truncation")
	}
	if len(fetchedTicket.ResolutionMessage) != len(longMsg) {
		t.Errorf("Expected length %d, got %d", len(longMsg), len(fetchedTicket.ResolutionMessage))
	}
}
