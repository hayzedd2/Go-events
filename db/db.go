package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
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
	 dateTime DATETIME NOT NULL,
	 location TEXT NOT NULL,
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
