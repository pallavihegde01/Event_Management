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

// Event represents the full event structure for frontend
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
	BookingLink string    `json:"bookingLink"`
	Source      string    `json:"source"`
	Verified    bool      `json:"verified"`
	IsPublic    bool      `json:"is_public"`
}

var db *sql.DB

func main() {
	var err error

	// Open SQLite database (creates file if not exists)
	db, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		log.Fatalf("❌ Failed to open database: %v\n", err)
	}
	defer db.Close()

	// Test DB connection
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Database connection failed: %v\n", err)
	}

	log.Println("✅ Database connected successfully")

	// Create table if not exists
	createTable()

	// Register HTTP handlers
	http.HandleFunc("/api/v1/events", corsMiddleware(createEvent))
	http.HandleFunc("/api/v1/events/", corsMiddleware(getEvent))

	log.Println("🚀 Server running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("❌ Server failed: %v\n", err)
	}
}

// Create table with all necessary fields
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
		log.Fatalf("❌ Failed to create table: %v\n", err)
	}
	log.Println("✅ Table checked/created")
}

// CORS middleware to allow frontend fetches
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// POST /api/v1/events
func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input Event
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Insert into database
	result, err := db.Exec(`
		INSERT INTO events
		(title, company, description, start_time, end_time, status, venue, organizers, booking_link, source, verified, is_public)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		input.Title,
		input.Company,
		input.Description,
		input.StartTime.Format(time.RFC3339),
		input.EndTime.Format(time.RFC3339),
		input.Status,
		input.Venue,
		input.Organizers,
		input.BookingLink,
		input.Source,
		input.Verified,
		input.IsPublic,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	input.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

// GET /api/v1/events/{id}
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

	var e Event
	var startStr, endStr string

	err = db.QueryRow(`
		SELECT id, title, company, description, start_time, end_time, status, venue, organizers, booking_link, source, verified, is_public
		FROM events
		WHERE id = ?`, id).Scan(
		&e.ID,
		&e.Title,
		&e.Company,
		&e.Description,
		&startStr,
		&endStr,
		&e.Status,
		&e.Venue,
		&e.Organizers,
		&e.BookingLink,
		&e.Source,
		&e.Verified,
		&e.IsPublic,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e.StartTime, _ = time.Parse(time.RFC3339, startStr)
	e.EndTime, _ = time.Parse(time.RFC3339, endStr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}