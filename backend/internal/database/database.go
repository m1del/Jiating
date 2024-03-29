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

	// done
	CreateAdmin(ctx context.Context, admin models.Admin) (string, error)

	GetAllAdmins(ctx context.Context, page, pageSize int) ([]models.Admin, error)
	GetAllAdminsExceptFounder(ctx context.Context, page, pageSize int) ([]models.Admin, error)
	GetAdmin(ctx context.Context, field, value string) (*models.Admin, error)

	GetAdminCount(ctx context.Context) (int, error)

	UpdateAdmin(ctx context.Context, admin models.Admin) error
	// needs revision
	//TODO: deletion

	// event operations

	// done
	CreateEvent(ctx context.Context, event models.Event, adminID string) (string, error)

	// TODO: refactor
	// GetAuthorsByEventID(eventID string) ([]models.Admin, error)
	// CreateEvent(event models.Event, adminIDs []string) (string, error)
	// UpdateEvent(event models.Event, editorAdminID string, newImages []models.EventImage, removedImageIDs []string, newDisplayImageID string) error
	// UpdateEventByID(eventID string, req models.UpdateEventRequest) error
	// GetEventByID(eventID string) (*models.Event, error)
	// GetLastSevenPublishedEvents() ([]models.Event, error)

	// event helpers
	//UpdateDynamicEventFields(tx *sql.Tx, eventID string, updatedData map[string]interface{}) error
	//UpdateEventAuthorship(tx *sql.Tx, eventID, editorAdminID string) error

	// image operations
	//AddImageToEvent(image models.EventImage, eventID string) error
	//RemoveImageFromEvent(imageID string) error
	//SetDisplayImageForEvent(imageID string, eventID string) error
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

func New(db *sql.DB) Service {
	if db == nil {
		// Create a real database connection if db is nil
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
		var err error
		db, err = sql.Open("pgx", connStr)
		if err != nil {
			loggers.Error.Fatalf("error connecting to the database: %v", err)
		}

		// Initialize tables
		loggers.Debug.Println("initializing tables...")
		if err := initTables(db); err != nil {
			loggers.Error.Fatalf("error initializing tables: %v", err)
		}
	}

	return &service{db: db}
}

func initTables(db *sql.DB) error {
	if err := createAdminTable(db); err != nil {
		loggers.Error.Fatalf("error creating admins table: %v", err)
	}

	if err := createEventTable(db); err != nil {
		loggers.Error.Fatalf("error creating events table: %v", err)
	}

	if err := createEventAuthorTable(db); err != nil {
		loggers.Error.Fatalf("error creating event authors table: %v", err)
	}

	if err := createImageTable(db); err != nil {
		loggers.Error.Fatalf("error creating images table: %v", err)
	}

	return nil
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
