package database

import (
	"backend/internal/models"
	"backend/loggers"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// ===== internal ===== //

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
    );
	
	CREATE INDEX IF NOT EXISTS idx_admins_email on admins (email);`

	_, err := db.Exec(createAdminTableSQL)
	if err != nil {
		loggers.Error.Printf("Error creating admin table: %v", err)
		return err
	}

	// insert default (permanent) admin with current timestamp for created_at and updated_at
	const insertAdmin = `
    INSERT INTO admins (name, email, position, status, created_at, updated_at)
    VALUES ('Jiating', 'jiating.lion.dragon@gmail.com', 'Founder', 'permanent', NOW(), NOW())
    ON CONFLICT (email) DO NOTHING;`

	_, err = db.Exec(insertAdmin)
	if err != nil {
		loggers.Error.Printf("Error inserting default admin: %v", err)
		return err
	}

	return nil
}

// ===== external ===== //

func (s *service) GetAllAdmins() ([]models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, 
	position, status 
	FROM admins
	WHERE deleted_at IS NULL`
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

func (s *service) CreateAdmin(admin models.Admin) (string, error) {
	// validate admin email format and check if it already exists
	if !isValidEmail(admin.Email) {
		return "", errors.New("invalid email format")
	}

	exists, err := s.emailExists(admin.Email)
	if err != nil {
		loggers.Error.Printf("Error checking if email exists: %v", err)
		return "", err
	}
	if exists {
		return "", errors.New("email already exists")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var id string
	const query = `INSERT INTO admins (
        created_at, updated_at, name, email, position, status
    ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	currTime := time.Now()
	err = s.db.QueryRowContext(
		ctx, query, currTime, currTime, admin.Name, admin.Email,
		admin.Position, admin.Status,
	).Scan(&id)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505": // unique_violation
				return "", errors.New("email already exists")
			case "23503": // foreign_key_violation
				return "", errors.New("foreign key violation")
			default:
				return "", errors.New("database error: " + err.Code.Name())
			}
		}
		loggers.Error.Printf("Error creating admin: %v", err)
		return "", err
	}

	return id, nil
}

// func (s *service) AssociateAdminWithEvent(ctx context.Context, adminID, eventID string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const query = `INSERT INTO event_authors (admin_id, event_id) VALUES ($1, $2)`

// 	_, err := s.db.ExecContext(ctx, query, adminID, eventID)
// 	if err != nil {
// 		if err, ok := err.(*pq.Error); ok {
// 			switch err.Code {
// 			case "23503": // foreign_key_violation
// 				return errors.New("invalid adminID or eventID")
// 			default:
// 				return errors.New("unknown error: " + err.Code.Name())
// 			}
// 		}
// 		loggers.Error.Printf("Error associating admin with event: %v", err)
// 		return err
// 	}

// 	return nil
// }

func (s *service) AssociateAdminWithEventTx(ctx context.Context, tx *sql.Tx, adminID, eventID string) error {
	const query = `INSERT INTO event_authors (admin_id, event_id) VALUES ($1, $2)`

	_, err := tx.ExecContext(ctx, query, adminID, eventID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23503": // foreign_key_violation
				return errors.New("invalid adminID or eventID")
			default:
				return errors.New("unknown error: " + err.Code.Name())
			}
		}
		loggers.Error.Printf("Error associating admin with event: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteAdminByID(adminID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const softDeleteAdminQuery = `
	UPDATE admins 
	SET deleted_at = $1 
	WHERE id = $2 
		AND status != 'permanent' 
	RETURNING id`

	var deletedAdminID string
	err := s.db.QueryRowContext(ctx, softDeleteAdminQuery, time.Now(), adminID).Scan(&deletedAdminID)
	if err != nil {
		if err == sql.ErrNoRows {
			// either the admin is permanent or does not exist
			return fmt.Errorf("cannot delete admin: admin is either permanent or does not exist")
		}
		loggers.Error.Printf("Error deleting admin: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteAdminByEmail(adminEmail string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const softDeleteAdminQuery = `
	UPDATE admins
	SET deleted_at = $1
	WHERE email = $2
		AND status != 'permanent'
	RETURNING id`

	var deletedAdminID string
	err := s.db.QueryRowContext(ctx, softDeleteAdminQuery, time.Now(), adminEmail).Scan(&deletedAdminID)
	if err != nil {
		if err == sql.ErrNoRows {
			// either the admin is permanent or does not exist
			return fmt.Errorf("cannot delete admin: admin is either permanent or does not exist")
		}
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
    FROM admins 
	WHERE id = $1 
		AND deleted_at IS NULL`
	row := s.db.QueryRowContext(ctx, query, adminID)

	var admin models.Admin
	err := row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
		&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			loggers.Error.Printf("No admin found with ID: %v", adminID)
			return nil, fmt.Errorf("admin not found")
		}
		loggers.Error.Printf("Error getting admin: %v", err)
		return nil, err
	}

	return &admin, nil
}

func (s *service) GetAdminByEmail(adminEmail string) (*models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, position, status 
	FROM admins 
	WHERE email = $1
		AND deleted_at IS NULL`
	row := s.db.QueryRowContext(ctx, query, adminEmail)

	var admin models.Admin
	err := row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
		&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			loggers.Error.Printf("No admin found with email: %v", adminEmail)
			return nil, err
		}
		loggers.Error.Printf("Error getting admin: %v", err)
		return nil, err
	}

	return &admin, nil
}

func (s *service) UpdateAdmin(adminID string, updateData models.AdminUpdateData) error {
	// TODO add handler
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
		return fmt.Errorf("cannot update a permanent admin")
	}

	// Uupdate logic
	const updateQuery = `
    UPDATE admins
    SET name = $1, position = $2, status = $3, email = $4, updated_at = $5
    WHERE id = $6`
	_, err = s.db.ExecContext(ctx, updateQuery, updateData.Name, updateData.Position, updateData.Status, updateData.Email, time.Now(), adminID)
	if err != nil {
		loggers.Error.Printf("Error updating admin: %v", err)
		return err
	}

	return nil
}

func (s *service) GetAdminCount() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const query = `SELECT COUNT(*) FROM admins`
	row := s.db.QueryRowContext(ctx, query)

	var count int
	err := row.Scan(&count)
	if err != nil {
		loggers.Error.Printf("Error getting admin count: %v", err)
		return 0, err
	}

	return count, nil
}
