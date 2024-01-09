package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"time"
)

func createEventTable(db *sql.DB) error {
	createEventTableSQL := `
    CREATE TABLE IF NOT EXISTS events (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
        event_name VARCHAR(255) NOT NULL,
        date VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        content TEXT NOT NULL,
        is_draft BOOLEAN NOT NULL,
        published_at TIMESTAMP WITH TIME ZONE
    );`

	_, err := db.Exec(createEventTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating event table: %v", err)
		return err
	}

	return nil
}

func createEventAuthorTable(db *sql.DB) error {
	createEventAuthorTableSQL := `
    CREATE TABLE IF NOT EXISTS event_authors (
        admin_id UUID NOT NULL,
        event_id UUID NOT NULL,
        PRIMARY KEY (admin_id, event_id),
        FOREIGN KEY (admin_id) REFERENCES admins(id) ON DELETE CASCADE,
        FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
    );`

	_, err := db.Exec(createEventAuthorTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating event_author table: %v", err)
		return err
	}

	return nil
}

func (s *service) CreateEvent(event models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	publishedTime := time.Time{}

	if event.IsDraft == false {
		publishedTime = time.Now()
	}

	const query = `INSERT INTO events (
        id, created_at, updated_at, admin_id, event_name, date, description, content, is_draft, published_at, image_id
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := s.db.ExecContext(
		ctx, query, event.ID, time.Now(), time.Now(), event.AdminID, event.EventName, event.Date, event.Description, event.Content, event.IsDraft, publishedTime, event.ImageID,
	)
	if err != nil {
		loggers.Error.Printf("error creating event: %v", err)
		return err
	}

	return nil
}

func (s *service) UpdateEvent(event models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	publishedTime := time.Time{}

	if event.IsDraft == false {
		publishedTime = time.Now()
	}

	const query = `UPDATE events SET
		updated_at = $1,
		admin_id = $2,
		event_name = $3,
		date = $4,
		description = $5,
		content = $6,
		is_draft = $7,
		published_at = $8,
		image_id = $9
		WHERE id = $10`

	_, err := s.db.ExecContext(
		ctx, query, time.Now(), event.AdminID, event.EventName, event.Date, event.Description, event.Content, event.IsDraft, publishedTime, event.ImageID, event.ID,
	)
	if err != nil {
		loggers.Error.Printf("error updating event: %v", err)
		return err
	}

	return nil
}

func (s *service) GetEventByID(id string) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	loggers.Debug.Println("Querying events table...")
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		loggers.Error.Printf("Error querying events table: %v", err)
		return nil, err
	}
	defer rows.Close()

	event := new(models.Event)
	if rows.Next() {
		err = rows.Scan(
			&event.ID, &event.CreatedAt, &event.UpdatedAt, &event.AdminID, &event.EventName, &event.Date, &event.Description, &event.Content, &event.IsDraft, &event.PublishedAt, &event.ImageID,
		)
		if err != nil {
			loggers.Error.Printf("Error scanning event row: %v", err)
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows // If no rows are found
	}

	return event, nil
}
