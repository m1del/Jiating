package database

import "regexp"

// data validation and sanitization

// isValidEmail checks if the email is in a valid format.
func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

// emailExists checks if the email already exists in the database.
func (s *service) emailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)`
	err := s.db.QueryRow(query, email).Scan(&exists)
	return exists, err
}
