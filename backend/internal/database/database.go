package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	GetAllAdmins() ([]*models.Admin, error)
	CreateAdmin(admin models.Admin) error
	CreateEvent(event models.Event) error
	UpdateEvent(event models.Event) error
	GetEventByID(id string) (*models.Event, error)
}

type service struct {
	db *sql.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		loggers.Error.Fatalf("error connecting to the database: %v", err)
	}

	// initialize tables
	if err := createAdminTable(db); err != nil {
		loggers.Error.Fatalf("error creating admins table: %v", err)
	}

	if err := createImageTable(db); err != nil {
		loggers.Error.Fatalf("error creating images table: %v", err)
	}

	if err := createEventTable(db); err != nil {
		loggers.Error.Fatalf("error creating events table: %v", err)
	}

	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		loggers.Error.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

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
