package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

// data validation
func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// status codes

func determineEmailStatusCode(err error) int {
	switch err.Error() {
	case "invalid email format", "input too long", "invalid admin status":
		return http.StatusBadRequest
	case "email already exists":
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
