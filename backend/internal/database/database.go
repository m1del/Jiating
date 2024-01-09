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

	// admin operations
	GetAllAdmins() ([]models.Admin, error)
	GetAdminByID(adminID string) (*models.Admin, error)
	GetAllAdminsExceptFounder() ([]models.Admin, error)
	CreateAdmin(admin models.Admin) error
	DeleteAdmin(adminID string) error
	UpdateAdmin(admin models.Admin) error
	AssociateAdminWithEvent(adminID string, eventID string) error

	// event operations
	GetAuthorsByEventID(eventID string) ([]models.Admin, error)
	CreateEvent(event models.Event, adminIDs []string) error
	UpdateEvent(event models.Event, editorAdminID string) error
	GetEventByID(eventID string) (*models.Event, error)
	GetLastSevenPublishedEvents() ([]models.Event, error)

	// image operations
	AddImageToEvent(image models.EventImage, eventID string) error
	RemoveImageFromEvent(imageID string) error
	SetDisplayImageForEvent(imageID string, eventID string) error
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
	if err := initTables(db); err != nil {
		loggers.Error.Fatalf("error initializing tables: %v", err)
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

func initTables(db *sql.DB) error {
	if err := createAdminTable(db); err != nil {
		loggers.Error.Fatalf("error creating admins table: %v", err)
	}

	if err := createImageTable(db); err != nil {
		loggers.Error.Fatalf("error creating images table: %v", err)
	}

	if err := createEventTable(db); err != nil {
		loggers.Error.Fatalf("error creating events table: %v", err)
	}

	if err := createEventAuthorTable(db); err != nil {
		loggers.Error.Fatalf("error creating event authors table: %v", err)
	}

	return nil
}
