package database

import (
	"backend/internal/models"
	"errors"
	"html"
	"regexp"
	"strings"
)

// pagination utils
func getOffset(page, pageSize int) int {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // default page size
	}
	return (page - 1) * pageSize
}

// data validation and sanitization

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func (s *service) emailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)`
	err := s.db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

func SanitizeAdminInput(admin *models.Admin) error {
	if admin == nil {
		return errors.New("admin is nil")
	}
	if !isValidEmail(admin.Email) {
		return errors.New("invalid email format")
	}
	// trim whitepaces and normalize to lowercase
	admin.Email = strings.ToLower(strings.TrimSpace(admin.Email))
	admin.Position = strings.TrimSpace(admin.Position)
	// trim whitespace for name, keep casing
	admin.Name = strings.TrimSpace(admin.Name)
	// length limits from database schema
	if len(admin.Name) > 255 || len(admin.Position) > 255 || len(admin.Email) > 255 {
		return errors.New("input too long")
	}
	// escaping special characters for name to prevent XSS attacks
	admin.Name = html.EscapeString(admin.Name)
	// validate admin status
	if admin.Status != "active" && admin.Status != "inactive" && admin.Status != "hiatus" {
		return errors.New("invalid admin status")
	}
	return nil
}
