package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"time"
)

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
