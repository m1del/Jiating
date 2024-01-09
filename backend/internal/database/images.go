package database

import (
	"backend/loggers"
	"database/sql"
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
