package database

import (
	"backend/loggers"
	"database/sql"
)

func createImageTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS images (
        id SERIAL PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        image_url VARCHAR(255) NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating admin table: %v", err)
		return err
	}

	return nil
}
