package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"time"
)

// ===== internal ===== //

func createEventTable(db *sql.DB) error {
	createEventTableSQL := `
    CREATE TABLE IF NOT EXISTS events (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
        event_title VARCHAR(255) NOT NULL,
		meta_title VARCHAR(255) NOT NULL,
		slug  VARCHAR(255) NOT NULL UNIQUE,
        date TIMESTAMP WITH TIME ZONE NOT NULL,
        description VARCHAR(2000) NOT NULL,
        content TEXT NOT NULL,
        is_draft BOOLEAN NOT NULL,
        published_at TIMESTAMP WITH TIME ZONE
	);
	
	CREATE INDEX IF NOT EXISTS idx_events_slug ON events(slug);`

	_, err := db.Exec(createEventTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating event table: %v", err)
		return err
	}

	return nil
}

// ===== exported ===== //

// CRUD operations

// Create operations:

func (s *service) CreateEvent(ctx context.Context, event models.Event, adminID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// start a transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		loggers.Error.Printf("Error starting transaction: %v", err)
		return "", err
	}

	// if the image is a draft, set the published_at field
	currentTime := time.Now()
	if !event.IsDraft {
		event.PublishedAt = &currentTime
	} else {
		event.PublishedAt = nil
	}

	// insert event data into events table
	var eventID string
	eventInsertQuery := `
	INSERT INTO EVENTS (
		id, created_at, updated_at, 
		event_title, meta_title, slug, date, 
		description, content, is_draft, published_at
	) VALUES (
		$1, $2, $3,
		$4, $5, $6, $7,
		$8, $9, $10, $11
	) RETURNING id`
	if err := tx.QueryRowContext(
		ctx, eventInsertQuery, event.ID, currentTime, currentTime,
		event.EventTitle, event.Metatitle, event.Slug, event.Date,
		event.Description, event.Content, event.IsDraft, event.PublishedAt).
		Scan(&eventID); err != nil {
		tx.Rollback()
		loggers.Error.Printf("Error inserting event: %v", err)
		return "", err
	}

	// associate admiin as author of event
	s.associateAdminWithEventTx(ctx, tx, adminID, eventID) // error handling is done in the function

	// insert image metadata into images table
	for _, img := range event.Images {
		if err := s.AddImageToEventTx(tx, img, eventID); err != nil {
			tx.Rollback()
			loggers.Error.Printf("Error adding image to event: %v", err)
			return "", err
		}
	}

	// TODO: handle display image
	err = tx.Commit()
	if err != nil {
		loggers.Error.Printf("committing create event: %v", err)
		return "", err
	}
	return eventID, nil
}

// Read Operations:

// func (s *service) GetEventByID(eventID string) (*models.Event, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	// fetch the event details
// 	const getEventQuery = `SELECT id, created_at, updated_at, event_name, date, description, content, is_draft, published_at FROM events WHERE id = $1`
// 	row := s.db.QueryRowContext(ctx, getEventQuery, eventID)

// 	var event models.Event
// 	err := row.Scan(&event.ID, &event.CreatedAt, &event.UpdatedAt, &event.EventName, &event.Date, &event.Description, &event.Content, &event.IsDraft, &event.PublishedAt)
// 	if err != nil {
// 		loggers.Error.Printf("Error retrieving event: %v", err)
// 		return nil, err
// 	}

// 	// fetch associated images
// 	const getImagesQuery = `SELECT id, image_url, is_display FROM images WHERE event_id = $1`
// 	images, err := s.getImagesByEventID(ctx, eventID)
// 	if err != nil {
// 		loggers.Error.Printf("Error retrieving images for event: %v", err)
// 		return nil, err
// 	}
// 	event.Images = images

// 	// fetch associated authors
// 	authors, err := s.GetAuthorsByEventID(eventID)
// 	if err != nil {
// 		loggers.Error.Printf("Error retrieving authors for event: %v", err)
// 		return nil, err
// 	}
// 	event.Authors = authors

// 	return &event, nil
// }

// func (s *service) GetAuthorsByEventID(eventID string) ([]models.Admin, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const query = `
//     SELECT a.id, a.created_at, a.updated_at, a.deleted_at, a.name, a.email, a.position, a.status
//     FROM admins a
//     INNER JOIN event_authors ea ON a.id = ea.admin_id
//     WHERE ea.event_id = $1`

// 	rows, err := s.db.QueryContext(ctx, query, eventID)
// 	if err != nil {
// 		loggers.Error.Printf("Error retrieving authors for event: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var authors []models.Admin
// 	for rows.Next() {
// 		var author models.Admin
// 		if err := rows.Scan(&author.ID, &author.CreatedAt, &author.UpdatedAt, &author.DeletedAt, &author.Name, &author.Email, &author.Position, &author.Status); err != nil {
// 			loggers.Error.Printf("Error scanning author: %v", err)
// 			continue
// 		}
// 		authors = append(authors, author)
// 	}

// 	if err := rows.Err(); err != nil {
// 		loggers.Error.Printf("Error iterating over authors: %v", err)
// 		return nil, err
// 	}

// 	return authors, nil
// }

// func (s *service) GetLastSevenPublishedEvents() ([]models.Event, error) {
// 	eventIDs, err := s.getLastSevenPublishedEventIDs()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var events []models.Event
// 	for _, id := range eventIDs {
// 		event, err := s.GetEventByID(id)
// 		if err != nil {
// 			// log and skip this event
// 			loggers.Error.Printf("Error retrieving event by ID: %v", err)
// 			continue
// 		}
// 		events = append(events, *event)
// 	}

// 	return events, nil
// }

// func (s *service) getLastSevenPublishedEventIDs() ([]string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const getEventIDsQuery = `
//     SELECT id FROM events
//     WHERE is_draft = false
//     ORDER BY published_at DESC
//     LIMIT 7`

// 	rows, err := s.db.QueryContext(ctx, getEventIDsQuery)
// 	if err != nil {
// 		loggers.Error.Printf("Error retrieving last seven published event IDs: %v", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var eventIDs []string
// 	for rows.Next() {
// 		var id string
// 		if err := rows.Scan(&id); err != nil {
// 			continue
// 		}
// 		eventIDs = append(eventIDs, id)
// 	}

// 	if err := rows.Err(); err != nil {
// 		loggers.Error.Printf("Error iterating over event IDs: %v", err)
// 		return nil, err
// 	}

// 	return eventIDs, nil
// }

// TODO: Update Operations

// TODO: Delete Operations

// func (s *service) UpdateEventByID(eventID string, req models.UpdateEventRequest) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	// start a transaction
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		loggers.Error.Printf("Error starting transaction: %v", err)
// 		return err
// 	}

// 	// build and execute the dynamic update query for the event
// 	if err := s.UpdateDynamicEventFields(tx, eventID, req.UpdatedData); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// add new images
// 	for _, img := range req.NewImages {
// 		if err := s.AddImageToEventTx(tx, img, eventID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error adding image to event: %v", err)
// 			return err
// 		}
// 	}

// 	// remove images
// 	for _, imgID := range req.RemovedImageIDs {
// 		if err := s.RemoveImageFromEventTx(tx, imgID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error removing image from event: %v", err)
// 			return err
// 		}
// 	}

// 	// set new display image if provided
// 	if req.NewDisplayImage != "" {
// 		if err := s.SetDisplayImageForEventTx(tx, req.NewDisplayImage, eventID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error setting display image: %v", err)
// 			return err
// 		}
// 	}

// 	// check and update the authorship
// 	if err := s.UpdateEventAuthorship(tx, eventID, req.EditorAdminID); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// commit the transaction
// 	if err := tx.Commit(); err != nil {
// 		loggers.Error.Printf("Error committing transaction: %v", err)
// 		return err
// 	}

// 	return nil
// }

// update event helper functions
// func (s *service) UpdateDynamicEventFields(tx *sql.Tx, eventID string, updatedData map[string]interface{}) error {
// 	updateQuery := "UPDATE events SET "
// 	var args []interface{}
// 	argID := 1

// 	for key, value := range updatedData {
// 		updateQuery += fmt.Sprintf("%s = $%d, ", key, argID)
// 		args = append(args, value)
// 		argID++
// 	}

// 	updateQuery += fmt.Sprintf("updated_at = $%d WHERE id = $%d", argID, argID+1)
// 	args = append(args, time.Now(), eventID)

// 	_, err := tx.Exec(updateQuery, args...)
// 	if err != nil {
// 		loggers.Error.Printf("Error updating event fields: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) UpdateEventAuthorship(tx *sql.Tx, eventID, editorAdminID string) error {
// 	// check if this admin is already an author of the event
// 	const checkAuthorQuery = `SELECT EXISTS (SELECT 1 FROM event_authors WHERE admin_id = $1 AND event_id = $2)`
// 	var exists bool
// 	if err := tx.QueryRow(checkAuthorQuery, editorAdminID, eventID).Scan(&exists); err != nil {
// 		loggers.Error.Printf("Error checking for existing author: %v", err)
// 		return err
// 	}

// 	// ff not already an author, add them
// 	if !exists {
// 		const addAuthorQuery = `INSERT INTO event_authors (admin_id, event_id) VALUES ($1, $2)`
// 		if _, err := tx.Exec(addAuthorQuery, editorAdminID, eventID); err != nil {
// 			loggers.Error.Printf("Error adding new event author: %v", err)
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (s *service) UpdateEvent(event models.Event, editorAdminID string, newImages []models.EventImage, removedImageIDs []string, newDisplayImageID string) error {
// 	//Note: this function is not used in the current implementation bc i cant figure it out lol
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	// start a transaction
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		loggers.Error.Printf("Error starting transaction: %v", err)
// 		return err
// 	}

// 	// update the event details
// 	const updateEventQuery = `UPDATE events SET event_name = $1, date = $2, description = $3, content = $4, is_draft = $5, published_at = $6, updated_at = $7 WHERE id = $8`
// 	_, err = tx.ExecContext(ctx, updateEventQuery, event.EventName, event.Date, event.Description, event.Content, event.IsDraft, event.PublishedAt, time.Now(), event.ID)
// 	if err != nil {
// 		tx.Rollback()
// 		loggers.Error.Printf("Error updating event: %v", err)
// 		return err
// 	}

// 	// add new images
// 	for _, img := range newImages {
// 		if err := s.AddImageToEventTx(tx, img, event.ID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error adding image to event: %v", err)
// 			return err
// 		}
// 	}

// 	// remove images
// 	for _, imgID := range removedImageIDs {
// 		if err := s.RemoveImageFromEvent(imgID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error removing image from event: %v", err)
// 			return err
// 		}
// 	}

// 	// set new display image if provided
// 	if newDisplayImageID != "" {
// 		if err := s.SetDisplayImageForEventTx(tx, newDisplayImageID, event.ID); err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error setting display image: %v", err)
// 			return err
// 		}
// 	}

// 	// check if this admin is already an author of the event
// 	const checkAuthorQuery = `SELECT EXISTS (SELECT 1 FROM event_authors WHERE admin_id = $1 AND event_id = $2)`
// 	var exists bool
// 	err = tx.QueryRowContext(ctx, checkAuthorQuery, editorAdminID, event.ID).Scan(&exists)
// 	if err != nil {
// 		tx.Rollback()
// 		loggers.Error.Printf("Error checking for existing author: %v", err)
// 		return err
// 	}

// 	// if not already an author, add them
// 	if !exists {
// 		const addAuthorQuery = `INSERT INTO event_authors (admin_id, event_id) VALUES ($1, $2)`
// 		_, err = tx.ExecContext(ctx, addAuthorQuery, editorAdminID, event.ID)
// 		if err != nil {
// 			tx.Rollback()
// 			loggers.Error.Printf("Error adding new event author: %v", err)
// 			return err
// 		}
// 	}

// 	// commit the transaction
// 	if err = tx.Commit(); err != nil {
// 		loggers.Error.Printf("Error committing transaction: %v", err)
// 		return err
// 	}

// 	return nil
// }
