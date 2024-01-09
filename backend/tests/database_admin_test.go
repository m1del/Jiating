package tests

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"
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

func TestDefaultAdminExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	expectedAdmin := models.Admin{
		Name:     "Jiating",
		Email:    "jiating.lion.dragon@gmail.com",
		Position: "Founder",
		Status:   "permanent",
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"}).
		AddRow("some-uuid", time.Now(), time.Now(), nil,
			expectedAdmin.Name, expectedAdmin.Email, expectedAdmin.Position, expectedAdmin.Status)

	mock.ExpectQuery("^SELECT (.+) FROM admins WHERE email = \\$1$").
		WithArgs(expectedAdmin.Email).
		WillReturnRows(rows)

	admin, err := s.GetAdminByEmail(expectedAdmin.Email)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, admin)
	assert.Equal(t, expectedAdmin.Email, admin.Email)
	assert.Equal(t, expectedAdmin.Name, admin.Name)
	assert.Equal(t, expectedAdmin.Position, admin.Position)
	assert.Equal(t, expectedAdmin.Status, admin.Status)

	// check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAdminByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	// Test data
	testEmail := "test@example.com"
	testAdmin := models.Admin{
		ID:        "test-uuid",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Test Admin",
		Email:     testEmail,
		Position:  "Test Position",
		Status:    "Active",
	}

	// Setup mock expectation
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"}).
		AddRow(testAdmin.ID, testAdmin.CreatedAt, testAdmin.UpdatedAt, testAdmin.DeletedAt, testAdmin.Name, testAdmin.Email, testAdmin.Position, testAdmin.Status)

	mock.ExpectQuery("^SELECT (.+) FROM admins WHERE email = \\$1$").
		WithArgs(testEmail).
		WillReturnRows(rows)

	// Call the function
	admin, err := s.GetAdminByEmail(testEmail)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, admin)
	assert.Equal(t, testAdmin, *admin)

	// check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCreateAdminSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	// Mock data
	admin := models.Admin{
		Name:     "john",
		Email:    "john@mail.com",
		Position: "idiot",
		Status:   "Active",
	}

	mock.ExpectExec("INSERT INTO admins").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), admin.Name, admin.Email, admin.Position, admin.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = s.CreateAdmin(admin)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCreateAdminFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	// Mock data
	admin := models.Admin{
		Name:     "tuan",
		Email:    "tuan@mail.com",
		Position: "dev",
		Status:   "Inactive",
	}

	mock.ExpectExec("INSERT INTO admins").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), admin.Name, admin.Email, admin.Position, admin.Status).
		WillReturnError(fmt.Errorf("sql error"))

	err = s.CreateAdmin(admin)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
