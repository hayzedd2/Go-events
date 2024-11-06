package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func getDatabasePath() string {
	if os.Getenv("ENVIRONMENT") == "production" {
		// Ensure the data directory exists
		err := os.MkdirAll("./data", 0755)
		if err != nil {
			log.Fatal("Could not create data directory:", err)
		}
		return filepath.Join("./data", "api.db")
	}
	return "api.db"
}
func InitDB() {
	var err error
	dbPath := getDatabasePath()
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not ping database:", err)
	}

	log.Printf("Connected to database at %s", dbPath)
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	userName TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	userId TEXT NOT NULL UNIQUE
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 startDate DATETIME NOT NULL,
	 startTime TEXT NOT NULL,
	 location TEXT NOT NULL,
	 category TEXT NOT NULL,
	 user_id TEXT,
	 FOREIGN KEY (user_id) REFERENCES users(userId)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create event table: " + err.Error())
	}
	createBookingsTable := `
	CREATE TABLE IF NOT EXISTS bookings (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 event_id INTEGER,
	 user_id TEXT,
	 FOREIGN KEY (event_id) REFERENCES events(id),
	 FOREIGN KEY (user_id) REFERENCES users(userId)
	)
	`
	_, err = DB.Exec(createBookingsTable)
	if err != nil {
		panic("Could not create booking table: " + err.Error())
	}
}
