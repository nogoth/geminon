package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Ticket struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	AreaOfConcern     string    `json:"area_of_concern"`
	AgentName         string    `json:"agent_name"`
	ResolutionMessage string    `json:"resolution_message"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

var DB *sql.DB

func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tickets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'Open',
		area_of_concern TEXT,
		agent_name TEXT,
		resolution_message TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createTableSQL)
	return err
}

func CreateTicket(ticket *Ticket) error {
	query := `INSERT INTO tickets (title, description, status, area_of_concern, agent_name, resolution_message, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	now := time.Now()
	res, err := DB.Exec(query, ticket.Title, ticket.Description, ticket.Status, ticket.AreaOfConcern, ticket.AgentName, ticket.ResolutionMessage, now, now)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	ticket.ID = int(id)
	ticket.CreatedAt = now
	ticket.UpdatedAt = now
	return nil
}

func UpdateTicket(id int, ticket *Ticket) error {
	query := `UPDATE tickets SET title = ?, description = ?, status = ?, area_of_concern = ?, agent_name = ?, resolution_message = ?, updated_at = ? WHERE id = ?`
	now := time.Now()
	res, err := DB.Exec(query, ticket.Title, ticket.Description, ticket.Status, ticket.AreaOfConcern, ticket.AgentName, ticket.ResolutionMessage, now, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("ticket not found")
	}
	ticket.UpdatedAt = now
	return nil
}

func GetTicket(id int) (*Ticket, error) {
	row := DB.QueryRow("SELECT id, title, description, status, area_of_concern, agent_name, resolution_message, created_at, updated_at FROM tickets WHERE id = ?", id)
	var t Ticket
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AreaOfConcern, &t.AgentName, &t.ResolutionMessage, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ListTickets(excludeArea, excludeStatus string) ([]Ticket, error) {
	query := "SELECT id, title, description, status, area_of_concern, agent_name, resolution_message, created_at, updated_at FROM tickets WHERE 1=1"
	var args []interface{}

	if excludeArea != "" {
		query += " AND area_of_concern != ?"
		args = append(args, excludeArea)
	}
	if excludeStatus != "" {
		query += " AND status != ?"
		args = append(args, excludeStatus)
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []Ticket
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AreaOfConcern, &t.AgentName, &t.ResolutionMessage, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}
