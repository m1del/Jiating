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

func (s *service) associateAdminWithEventTx(ctx context.Context, tx *sql.Tx, adminID, eventID string) error {
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

// ===== external ===== //

// CRUD operations for admins

// ========== CREATE ========== //

func (s *service) CreateAdmin(ctx context.Context, admin models.Admin) (string, error) {
	// validate admin email format and check if it already exists
	err := SanitizeAdminInput(&admin)
	if err != nil {
		loggers.Debug.Printf("invalid admin input: %v", err)
		return "", err
	}

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

// ========== READ ========== //
func scanAdmins(rows *sql.Rows) ([]models.Admin, error) {
	var admins []models.Admin
	for rows.Next() {
		var admin models.Admin
		if err := rows.Scan(
			&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
			&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status); err != nil {
			loggers.Error.Printf("scanning admin: %v", err)
			continue // skip partial results
		}
		admins = append(admins, admin)
	}
	if err := rows.Err(); err != nil {
		loggers.Error.Printf("Error iterating over admins: %v", err)
		return nil, err
	}
	return admins, nil
}

func (s *service) getAdmin(ctx context.Context, query string, param string) (*models.Admin, error) {
	row := s.db.QueryRowContext(ctx, query, param)

	var admin models.Admin
	err := row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt,
		&admin.DeletedAt, &admin.Name, &admin.Email, &admin.Position, &admin.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			loggers.Error.Printf("No admin found with parameter: %v", param)
			return nil, fmt.Errorf("admin not found")
		}
		loggers.Error.Printf("Error getting admin: %v", err)
		return nil, err
	}
	return &admin, nil
}

// define a function type for fetching admins to use in generalizzed handler
type AdminFetchFunc func(ctx context.Context, page, pageSize int) ([]models.Admin, error)

func (s *service) GetAllAdmins(ctx context.Context, page, pageSize int) ([]models.Admin, error) {
	offset := getOffset(page, pageSize)
	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, 
	position, status 
	FROM admins
	WHERE deleted_at IS NULL
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2`
	rows, err := s.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		loggers.Error.Printf("getting admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	return scanAdmins(rows)
}

func (s *service) GetAllAdminsExceptFounder(ctx context.Context, page, pageSize int) ([]models.Admin, error) {
	offset := getOffset(page, pageSize)
	const query = `
	SELECT id, created_at, updated_at, deleted_at, name, email, position, status 
	FROM admins 
	WHERE position != 'Founder' AND deleted_at IS NULL
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2`
	rows, err := s.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		loggers.Error.Printf("getting admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	return scanAdmins(rows)
}

// fetch admin by field: email or id exclusively
func (s *service) GetAdmin(ctx context.Context, field, value string) (*models.Admin, error) {
	query := fmt.Sprintf(`
    SELECT id, created_at, updated_at, deleted_at, name, email, position, status 
    FROM admins 
    WHERE %s = $1 
    AND deleted_at IS NULL`, field)

	return s.getAdmin(ctx, query, value)
}

func (s *service) GetAdminCount(ctx context.Context) (int, error) {
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

// ========== UPDATE ========== //

func (s *service) UpdateAdmin(ctx context.Context, admin models.Admin) error {
	err := SanitizeAdminInput(&admin)
	if err != nil {
		loggers.Debug.Printf("invalid admin input: %v", err)
		return err
	}
	const query = `
	UPDATE admins SET
	updated_at = $1, name = $2, email = $3, position = $4, status = $5
	WHERE id = $6
	AND deleted_at IS NULL`

	_, err = s.db.ExecContext(
		ctx, query, time.Now(), admin.Name, admin.Email, admin.Position, admin.Status, admin.ID,
	)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505": // unique_violation
				return errors.New("email already exists")
			default:
				return errors.New("database error: " + err.Code.Name())
			}
		}
		loggers.Error.Printf("updated admin: %v", err)
		return err
	}
	return nil
}

// TODO: everything below this line needs to be refactored

// ========== DELETE ========== //

func (s *service) DeleteAdmin(ctx context.Context, param string) (models.Admin, error) {
	return models.Admin{}, nil
}

// func (s *service) DeleteAdminByID(adminID string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const softDeleteAdminQuery = `
// 	UPDATE admins
// 	SET deleted_at = $1
// 	WHERE id = $2
// 		AND status != 'permanent'
// 	RETURNING id`

// 	var deletedAdminID string
// 	err := s.db.QueryRowContext(ctx, softDeleteAdminQuery, time.Now(), adminID).Scan(&deletedAdminID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// either the admin is permanent or does not exist
// 			return fmt.Errorf("cannot delete admin: admin is either permanent or does not exist")
// 		}
// 		loggers.Error.Printf("Error deleting admin: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func (s *service) DeleteAdminByEmail(adminEmail string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	const softDeleteAdminQuery = `
// 	UPDATE admins
// 	SET deleted_at = $1
// 	WHERE email = $2
// 		AND status != 'permanent'
// 	RETURNING id`

// 	var deletedAdminID string
// 	err := s.db.QueryRowContext(ctx, softDeleteAdminQuery, time.Now(), adminEmail).Scan(&deletedAdminID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// either the admin is permanent or does not exist
// 			return fmt.Errorf("cannot delete admin: admin is either permanent or does not exist")
// 		}
// 		loggers.Error.Printf("Error deleting admin: %v", err)
// 		return err
// 	}

// 	return nil
// }
