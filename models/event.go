package models

import (
	"time"

	"example.com/go-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	eventID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = eventID

	return nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event

	for row.Next() {
		var event Event
		err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventsById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?
    `

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Ensure the statement is closed after use

	// Execute the statement with the provided parameters
	if _, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID); err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
    `

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Ensure the statement is closed after use

	// Execute the statement with the provided parameters
	if _, err = stmt.Exec(e.ID); err != nil {
		return err
	}

	return nil
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id) VALUES (?, ?)
    `

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Ensure the statement is closed after use

	// Execute the statement with the provided parameters
	if _, err = stmt.Exec(e.ID, userId); err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations WHERE event_id = ? AND user_id = ?
    `

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Ensure the statement is closed after use

	// Execute the statement with the provided parameters
	if _, err = stmt.Exec(e.ID, userId); err != nil {
		return err
	}

	return nil
}
