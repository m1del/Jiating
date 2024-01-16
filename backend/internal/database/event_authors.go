package database

import (
	"backend/loggers"
	"database/sql"
)

func createEventAuthorTable(db *sql.DB) error {
	createEventAuthorTableSQL := `
    CREATE TABLE IF NOT EXISTS event_authors (
        admin_id UUID NOT NULL,
        event_id UUID NOT NULL,
        PRIMARY KEY (admin_id, event_id),
        FOREIGN KEY (admin_id) REFERENCES admins(id) ON DELETE NO ACTION,
        FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
    );`

	_, err := db.Exec(createEventAuthorTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating event_author table: %v", err)
		return err
	}

	return nil
}

// CRUD OPERATIONS

// ========== CREATE ========== //

// ========== READ ========== //

// ========== UPDATE ========== //

// ========== DELETE ========== //
