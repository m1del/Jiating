package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"time"
)

func createImageTable(db *sql.DB) error {
	createImageTableSQL := `
    CREATE TABLE IF NOT EXISTS images (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
        image_url VARCHAR(255) NOT NULL,
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

func (s *service) AddImageToEvent(image models.EventImage, eventID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `INSERT INTO images (image_url, is_display, event_id) VALUES ($1, $2, $3)`

	_, err := s.db.ExecContext(ctx, query, image.ImageURL, image.IsDisplay, eventID)
	if err != nil {
		loggers.Error.Printf("Error adding image to event: %v", err)
		return err
	}

	return nil
}

func (s *service) RemoveImageFromEvent(imageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `DELETE FROM images WHERE id = $1`

	// todo: delete image from s3

	_, err := s.db.ExecContext(ctx, query, imageID)
	if err != nil {
		loggers.Error.Printf("Error removing image from event: %v", err)
		return err
	}

	return nil
}

func (s *service) SetDisplayImageForEvent(imageID, eventID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// start a transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		loggers.Error.Printf("Error starting transaction: %v", err)
		return err
	}

	// reset isDisplay for all images of this event
	const resetQuery = `UPDATE images SET is_display = false WHERE event_id = $1`
	_, err = tx.ExecContext(ctx, resetQuery, eventID)
	if err != nil {
		tx.Rollback()
		loggers.Error.Printf("Error resetting display images: %v", err)
		return err
	}

	// set isDisplay to true for the chosen image
	const updateQuery = `UPDATE images SET is_display = true WHERE id = $1`
	_, err = tx.ExecContext(ctx, updateQuery, imageID)
	if err != nil {
		tx.Rollback()
		loggers.Error.Printf("Error setting display image: %v", err)
		return err
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		loggers.Error.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil
}

func (s *service) getImagesByEventID(ctx context.Context, eventID string) ([]models.EventImage, error) {
	const query = `SELECT id, image_url, is_display FROM images WHERE event_id = $1`
	rows, err := s.db.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.EventImage
	for rows.Next() {
		var img models.EventImage
		if err := rows.Scan(&img.ID, &img.ImageURL, &img.IsDisplay); err != nil {
			continue
		}
		images = append(images, img)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return images, nil
}
