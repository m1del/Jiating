package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func createAdminTable(db *sql.DB) error {
	createAdminTableSQL := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE IF NOT EXISTS admins (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
        deleted_at TIMESTAMP WITH TIME ZONE,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        position VARCHAR(255) NOT NULL,
        status VARCHAR(50) NOT NULL
    );`

	_, err := db.Exec(createAdminTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating admin table: %v", err)
		return err
	}

	// insert default (permanent) admin
	const insertAdmin = `
	INSERT INTO admins (name, email, position, status)
	VALUES (Jiating, jiating.lion.dragon@gmail.com, Founder, permanent)
	ON CONFLICT (email) DO NOTHING;
	`

	_, err = db.Exec(insertAdmin)
	if err != nil {
		loggers.Error.Printf("Error inserting default admin: %v", err)
		return err
	}

	return nil
}

func (s *service) GetAllAdmins() ([]models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, 
	position, status 
	FROM admins`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		loggers.Error.Printf("Error getting admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	var admins []models.Admin
	for rows.Next() {
		var admin models.Admin
		if err := rows.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
			&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status); err != nil {
			loggers.Error.Printf("Error scanning admin: %v", err)
			continue
		}
		admins = append(admins, admin)
	}

	if err := rows.Err(); err != nil {
		loggers.Error.Printf("Error iterating over admins: %v", err)
		return nil, err
	}

	return admins, nil
}

func (s *service) GetAllAdminsExceptFounder() ([]models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, position, status 
	FROM admins WHERE position != 'Founder'`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		loggers.Error.Printf("Error getting admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	var admins []models.Admin
	for rows.Next() {
		var admin models.Admin
		if err := rows.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt, &admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status); err != nil {
			loggers.Error.Printf("Error scanning admin: %v", err)
			continue
		}
		admins = append(admins, admin)
	}

	if err := rows.Err(); err != nil {
		loggers.Error.Printf("Error iterating over admins: %v", err)
		return nil, err
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
		ctx, query, time.Now(), time.Now(), admin.Name, admin.Email,
		admin.Position, admin.Status,
	)
	if err != nil {
		loggers.Error.Printf("Error creating admin: %v", err)
		return err
	}

	return nil
}

func (s *service) AssociateAdminWithEvent(adminID, eventID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `INSERT INTO event_authors (admin_id, event_id) VALUES ($1, $2)`

	_, err := s.db.ExecContext(ctx, query, adminID, eventID)
	if err != nil {
		loggers.Error.Printf("Error associating admin with event: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteAdmin(adminID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// check if the admin is permanent
	var status string
	const getStatusQuery = `SELECT status FROM admins WHERE id = $1`
	err := s.db.QueryRowContext(ctx, getStatusQuery, adminID).Scan(&status)
	if err != nil {
		loggers.Error.Printf("Error retrieving admin status: %v", err)
		return err
	}

	if status == "permanent" {
		// return an error or handle the attempt to delete the permanent admin
		return fmt.Errorf("cannot delete a permanent admin")
	}

	// proceed with deletion logic if not permanent
	const deleteAdminQuery = `DELETE FROM admins WHERE id = $1`
	_, err = s.db.ExecContext(ctx, deleteAdminQuery, adminID)
	if err != nil {
		loggers.Error.Printf("Error deleting admin: %v", err)
		return err
	}

	return nil
}

func (s *service) GetAdminByID(adminID string) (*models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, position, status 
	FROM admins WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, adminID)

	var admin models.Admin
	err := row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
		&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status)
	if err != nil {
		loggers.Error.Printf("Error getting admin: %v", err)
		return nil, err
	}

	return &admin, nil
}

func (s *service) UpdateAdmin(admin models.Admin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
    UPDATE admins
    SET name = $1, email = $2, position = $3, status = $4, updated_at = $5
    WHERE id = $6`

	_, err := s.db.ExecContext(ctx, query, admin.Name, admin.Email,
		admin.Position, admin.Status, time.Now(), admin.ID)
	if err != nil {
		loggers.Error.Printf("Error updating admin: %v", err)
		return err
	}

	return nil
}
