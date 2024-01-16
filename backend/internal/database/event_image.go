package database

import (
	"backend/internal/models"
	"backend/loggers"
	"database/sql"
)

// Initialize image table in database on startup
func createImageTable(db *sql.DB) error {
	createImageTableSQL := `
    CREATE TABLE IF NOT EXISTS event_images (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        image_url VARCHAR(255) NOT NULL,
		alt_text VARCHAR(255) NOT NULL,
        is_display BOOLEAN NOT NULL,
        event_id UUID REFERENCES events(id)
    );`

	_, err := db.Exec(createImageTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating image table: %v", err)
		return err
	}

	return nil
}

// ========== CRUD OPERATIONS ========== //

// ========== CREATE ========== //
// Note: images aren't "created" per say
// image metadata is created when an event is created
// images are uploaded to s3 and the url is stored in the db

// addImageToEventTx adds an image to an event in a transaction
func (s *service) AddImageToEventTx(tx *sql.Tx, image models.EventImage, eventID string) error {
	const query = `
	INSERT INTO event_images(
		id, created_at, image_url, 
		alt_text, is_display, event_id
	) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := tx.Exec(
		query, image.ID, image.CreatedAt, image.ImageURL,
		image.AltText, image.IsDisplay, eventID)
	if err != nil {
		tx.Rollback()
		loggers.Error.Printf("Error adding image to event: %v", err)
		return err
	}
	return nil
}

// ========== READ ========== //

// ========== UPDATE ========== //

// ========== DELETE ========== //

// func (s *service) RemoveImageFromEvent(imageID string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const query = `DELETE FROM event_images WHERE id = $1`

// 	// todo: delete image from s3

// 	_, err := s.db.ExecContext(ctx, query, imageID)
// 	if err != nil {
// 		loggers.Error.Printf("Error removing image from event: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) RemoveImageFromEventTx(tx *sql.Tx, imageID string) error {
// 	const query = `DELETE FROM event_images WHERE id = $1`

// 	// todo: delete image from s3

// 	if _, err := tx.Exec(query, imageID); err != nil {
// 		loggers.Error.Printf("Error removing image from event: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) SetDisplayImageForEvent(imageID, eventID string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	// start a transaction
// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		loggers.Error.Printf("Error starting transaction: %v", err)
// 		return err
// 	}

// 	// reset isDisplay for all event_images of this event
// 	const resetQuery = `UPDATE event_images SET is_display = false WHERE event_id = $1`
// 	_, err = tx.ExecContext(ctx, resetQuery, eventID)
// 	if err != nil {
// 		tx.Rollback()
// 		loggers.Error.Printf("Error resetting display event_images: %v", err)
// 		return err
// 	}

// 	// set isDisplay to true for the chosen image
// 	const updateQuery = `UPDATE event_images SET is_display = true WHERE id = $1`
// 	_, err = tx.ExecContext(ctx, updateQuery, imageID)
// 	if err != nil {
// 		tx.Rollback()
// 		loggers.Error.Printf("Error setting display image: %v", err)
// 		return err
// 	}

// 	// commit the transaction
// 	if err = tx.Commit(); err != nil {
// 		loggers.Error.Printf("Error committing transaction: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) SetDisplayImageForEventTx(tx *sql.Tx, imageID, eventID string) error {

// 	// do nothing if imageID is empty
// 	if imageID == "" {
// 		return nil
// 	}

// 	// reset isDisplay for all event_images of this event
// 	const resetQuery = `UPDATE event_images SET is_display = false WHERE event_id = $1`
// 	if _, err := tx.Exec(resetQuery, eventID); err != nil {
// 		loggers.Error.Printf("Error resetting display event_images: %v", err)
// 		return err
// 	}

// 	// set isDisplay to true for the chosen image
// 	const updateQuery = `UPDATE event_images SET is_display = true WHERE id = $1`
// 	if _, err := tx.Exec(updateQuery, imageID); err != nil {
// 		loggers.Error.Printf("Error setting display image: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) getImagesByEventID(ctx context.Context, eventID string) ([]models.EventImage, error) {
// 	const query = `SELECT id, image_url, is_display FROM event_images WHERE event_id = $1`
// 	rows, err := s.db.QueryContext(ctx, query, eventID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var event_images []models.EventImage
// 	for rows.Next() {
// 		var img models.EventImage
// 		if err := rows.Scan(&img.ID, &img.ImageURL, &img.IsDisplay); err != nil {
// 			continue
// 		}
// 		event_images = append(event_images, img)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return event_images, nil
// }
