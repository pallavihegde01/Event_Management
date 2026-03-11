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
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Company     string    `json:"company"`
    Description string    `json:"description"`
    StartTime   time.Time `json:"start_time"`
    EndTime     time.Time `json:"end_time"`
    Status      string    `json:"status"`
    Venue       string    `json:"venue"`
    Organizers  string    `json:"organizers"`
    BookingLink string    `json:"booking_link"`
    Source      string    `json:"source"`
    Verified    bool      `json:"verified"`
    IsPublic    bool      `json:"is_public"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v\n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v\n", err)
	}

	log.Println("Database connected successfully")

	createTable()

	http.HandleFunc("/api/v1/events", corsMiddleware(eventsHandler))
	http.HandleFunc("/api/v1/events/", corsMiddleware(eventHandler))

	log.Println("Server running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createEvent(w, r)
		return
	}
	if r.Method == http.MethodGet {
		getAllEvents(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getEvent(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		deleteEvent(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var input Event
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`
	INSERT INTO events
	(title, company, description, start_time, end_time, status, venue, organizers, booking_link, source, verified, is_public)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		input.Title, input.Company, input.Description,
		input.StartTime.Format(time.RFC3339),
		input.EndTime.Format(time.RFC3339),
		input.Status, input.Venue, input.Organizers,
		input.BookingLink, input.Source, input.Verified, input.IsPublic,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	input.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
	SELECT id, title, company, description, start_time, end_time, status, venue, organizers, booking_link, source, verified, is_public
	FROM events`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		var startStr, endStr string
		rows.Scan(
			&e.ID, &e.Title, &e.Company, &e.Description,
			&startStr, &endStr,
			&e.Status, &e.Venue, &e.Organizers,
			&e.BookingLink, &e.Source, &e.Verified, &e.IsPublic,
		)
		e.StartTime, _ = time.Parse(time.RFC3339, startStr)
		e.EndTime, _ = time.Parse(time.RFC3339, endStr)
		events = append(events, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var e Event
	var startStr, endStr string
	err = db.QueryRow(`
	SELECT id, title, company, description, start_time, end_time, status, venue, organizers, booking_link, source, verified, is_public
	FROM events WHERE id = ?`, id).
		Scan(
			&e.ID, &e.Title, &e.Company, &e.Description,
			&startStr, &endStr,
			&e.Status, &e.Venue, &e.Organizers,
			&e.BookingLink, &e.Source, &e.Verified, &e.IsPublic,
		)

	if err == sql.ErrNoRows {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	e.StartTime, _ = time.Parse(time.RFC3339, startStr)
	e.EndTime, _ = time.Parse(time.RFC3339, endStr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event deleted"))
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		company TEXT,
		description TEXT,
		start_time TEXT,
		end_time TEXT,
		status TEXT,
		venue TEXT,
		organizers TEXT,
		booking_link TEXT,
		source TEXT,
		verified BOOLEAN,
		is_public BOOLEAN
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to create table: %v\n", err)
	}
	log.Println("Table checked/created")
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}