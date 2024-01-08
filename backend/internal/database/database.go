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

func createAdminTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS admins (
        id SERIAL PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        deleted_at TIMESTAMP NULL,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        position VARCHAR(255) NOT NULL,
        status VARCHAR(50) NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating admin table: %v", err)
		return err
	}

	return nil
}

func (s *service) GetAllAdmins() ([]*models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	loggers.Debug.Println("Querying admins table...")
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM admins")
	if err != nil {
		loggers.Error.Printf("Error querying admins table: %v", err)
		return nil, err
	}
	defer rows.Close()

	admins := make([]*models.Admin, 0)
	for rows.Next() {
		admin := new(models.Admin)
		err := rows.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt, &admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status)
		if err != nil {
			loggers.Error.Printf("Error scanning admin row: %v", err)
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func (s *service) CreateAdmin(admin models.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `INSERT INTO admins (
        created_at, updated_at, name, email, position, status
    ) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.ExecContext(
		ctx, query, time.Now(), time.Now(), admin.Name, admin.Email, admin.Position, admin.Status,
	)
	if err != nil {
		loggers.Error.Printf("Error creating admin: %v", err)
		return err
	}

	return nil
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

func createEventTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS events (
        id VARCHAR(36) PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
		admin_id INT REFERENCES admins(id),
        event_name VARCHAR(255) NOT NULL,
		date VARCHAR(10) NOT NULL,
		description TEXT,
		content TEXT, 
        is_draft BOOLEAN NOT NULL,
		published_at TIMESTAMP NULL,
		image_id INT REFERENCES images(id)
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating event table: %v", err)
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

func (s* service) UpdateEvent(event models.Event) error {
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

func (s* service) GetEventByID(id string) (*models.Event, error) {
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

