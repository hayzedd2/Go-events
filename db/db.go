package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	var connURL = os.Getenv("DATABASE_URL")
	if connURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	DB, err = sql.Open("postgres", connURL)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not ping database:", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Test the connection

	log.Printf("Connected to database.")
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	userName TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	userId TEXT NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	 id SERIAL PRIMARY KEY,
	 name TEXT NOT NULL,
	 description TEXT NOT NULL,
	 startDate TIMESTAMP NOT NULL,
	 startTime TEXT NOT NULL,
	 location TEXT NOT NULL,
	 category TEXT NOT NULL,
	 user_id TEXT,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 FOREIGN KEY (user_id) REFERENCES users(userId)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create event table: " + err.Error())
	}
	createBookingsTable := `
	CREATE TABLE IF NOT EXISTS bookings (
	 id SERIAL PRIMARY KEY,
	 event_id INTEGER,
	 user_id TEXT,
	 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	 FOREIGN KEY (event_id) REFERENCES events(id),
	 FOREIGN KEY (user_id) REFERENCES users(userId)
	)
	`
	_, err = DB.Exec(createBookingsTable)
	if err != nil {
		panic("Could not create booking table: " + err.Error())
	}
}
