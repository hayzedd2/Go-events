package models

import (
	"github.com/hayzedd2/Go-events/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	Location    string    `binding:"required"`
	UserId      string
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, dateTime, location, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
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
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.Location, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	defer rows.Close()
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.Location, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, dateTime= ?, location = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) DELETE() error {
	query := "DELETE FROM events WHERE id = ?"
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
	query := "INSERT INTO bookings(event_id, user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}

func (e Event) CancelBooking(userId string) error {
	query := "DELETE FROM bookings WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}
