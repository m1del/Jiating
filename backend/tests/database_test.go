package tests

import (
	"backend/internal/database"
	"backend/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAdmins(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	// se fixed timestamps for consistency in tests
	fixedTime := time.Now()

	// Mock data
	admins := []models.Admin{
		{
			ID:        "some-uuid-1",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
			Name:      "john",
			Email:     "john@mail.com",
			Position:  "backend developer",
			Status:    "Active",
		},
		{
			ID:        "some-uuid-2",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
			Name:      "tuan",
			Email:     "tuan@mail.com",
			Position:  "frontend developer",
			Status:    "Active",
		},
		{
			ID:        "some-uuid-3",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "justin",
			Email:     "justin@mail.com",
			Position:  "hospitalized",
			Status:    "Inactive",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"})
	for _, admin := range admins {
		rows.AddRow(admin.ID, admin.CreatedAt, admin.UpdatedAt, admin.DeletedAt, admin.Name, admin.Email, admin.Position, admin.Status)
	}

	mock.ExpectQuery("^SELECT (.+) FROM admins$").WillReturnRows(rows)

	resultAdmins, err := s.GetAllAdmins()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, resultAdmins)
	assert.Equal(t, admins, resultAdmins)

	// check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
