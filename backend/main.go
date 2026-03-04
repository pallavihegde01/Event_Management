package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	IsPublic  bool      `json:"is_public"`
}

var db *sql.DB

func main() {
	log.Println("MAIN FUNCTION STARTED")

	var err error

	db, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		log.Fatal("Database open error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	createTable()

	http.HandleFunc("/api/v1/events", createEvent)
	http.HandleFunc("/api/v1/events/", getEvent)

	log.Println("Server running on http://localhost:8081")

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		type TEXT,
		start_time TEXT,
		end_time TEXT,
		is_public BOOLEAN
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Table creation error:", err)
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Title     string `json:"title"`
		Type      string `json:"type"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		IsPublic  bool   `json:"is_public"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	start, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		http.Error(w, "Invalid start_time format", http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		http.Error(w, "Invalid end_time format", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		`INSERT INTO events (title, type, start_time, end_time, is_public)
		 VALUES (?, ?, ?, ?, ?)`,
		input.Title,
		input.Type,
		start.Format(time.RFC3339),
		end.Format(time.RFC3339),
		input.IsPublic,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()

	event := Event{
		ID:        int(id),
		Title:     input.Title,
		Type:      input.Type,
		StartTime: start,
		EndTime:   end,
		IsPublic:  input.IsPublic,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var event Event
	var startStr, endStr string

	err = db.QueryRow(
		`SELECT id, title, type, start_time, end_time, is_public
		 FROM events WHERE id = ?`, id).
		Scan(&event.ID, &event.Title, &event.Type, &startStr, &endStr, &event.IsPublic)

	if err == sql.ErrNoRows {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	event.StartTime, _ = time.Parse(time.RFC3339, startStr)
	event.EndTime, _ = time.Parse(time.RFC3339, endStr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}