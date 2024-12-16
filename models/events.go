package models

import (
	"errors"
	"github.com/hayzedd2/eventease-be/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Category    string    `binding:"required"`
	StartDate   time.Time `binding:"required"`
	StartTime   string    `binding:"required"`
	UserId      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookStruct struct {
	ID      int64
	EventId int64
	UserId  string
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, startDate, startTime, location, category, user_id) 
	VALUES ($1, $2, $3, $4,$5,$6,$7)
	RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var id int64
	err = stmt.QueryRow(e.Name, e.Description, e.StartDate, e.StartTime, e.Location, e.Category, e.UserId).Scan(&id)
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}
func GetAllEvents() ([]Event, error) {
	query := "SELECT * from events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.StartDate, &event.StartTime, &event.Location, &event.Category, &event.UserId, &event.CreatedAt, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.StartDate, &event.StartTime, &event.Location, &event.Category, &event.UserId,&event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &event, nil
}


func (e Event) Update() error {
	query := `
    UPDATE events 
    SET name = $1, description = $2, startDate = $3, startTime = $4, 
        location = $5, category = $6, updated_at = NOW()
    WHERE id = $7`

    _, err := db.DB.Exec(query, e.Name, e.Description, e.StartDate, 
        e.StartTime, e.Location, e.Category, e.ID)
    return err
}

func (e Event) DELETE() error {
	query := "DELETE FROM events WHERE id = $1"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}
	return nil

}

func (e Event) Book(userId string) error {
	query := "INSERT INTO bookings(event_id, user_id) VALUES ($1,$2)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}

func (e Event) CancelBooking(userId string) error {
	query := "DELETE FROM bookings WHERE event_id = $1 AND user_id = $2"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}

func GetAllBookings() ([]BookStruct, error) {
	query := "SELECT * from bookings"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var bookings []BookStruct
	for rows.Next() {
		var booking BookStruct
		err := rows.Scan(&booking.ID, &booking.EventId, &booking.UserId)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	defer rows.Close()
	return bookings, nil
}

func IsBooked(userId string, eventId int64) (bool, error) {
	query := `SELECT COUNT(*) FROM bookings WHERE user_id = $1 AND event_id = $2`
	var count int
	err := db.DB.QueryRow(query, userId, eventId).Scan(&count)
	if err != nil {
		return false, errors.New("error checking booking status")
	}
	return count > 0, nil
}
