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

func TestGetAdminByEmailQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminEmail := "error@example.com"

	// Simulate a database error during querying
	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE email = \\$1").
		WithArgs(adminEmail).WillReturnError(fmt.Errorf("database error"))

	// Call the function under test
	admin, err := s.GetAdminByEmail(adminEmail)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, admin)

	// Check if all expectations were met
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

	expectedID := "some-uuid"

	rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
	mock.ExpectQuery("INSERT INTO admins").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), admin.Name, admin.Email,
		admin.Position, admin.Status).
		WillReturnRows(rows)

	id, err := s.CreateAdmin(admin)

	assert.NoError(t, err)
	assert.Equal(t, expectedID, id, "Expected ID does not match returned ID")

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

	admin := models.Admin{
		Name:     "tuan",
		Email:    "tuan@mail.com",
		Position: "dev",
		Status:   "Inactive",
	}

	// Sset up the mock to expect a QueryRow and return an error
	mock.ExpectQuery(
		"INSERT INTO admins").WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), admin.Name, admin.Email,
		admin.Position, admin.Status).
		WillReturnError(fmt.Errorf("sql error"))

	id, err := s.CreateAdmin(admin)

	assert.Error(t, err)
	assert.Empty(t, id, "Expected ID to be empty when an error occurs")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAdminByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminID := "some-uuid"
	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("Active")

	// mock query to check admin status
	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)

	// mock successful execution of the delete query
	mock.ExpectExec("DELETE FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnResult(sqlmock.NewResult(1, 1))

	// call the function under test
	err = s.DeleteAdminByID(adminID)

	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAdminByIDPermanent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminID := "some-uuid"
	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("permanent")

	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)
	err = s.DeleteAdminByID(adminID)

	assert.Error(t, err)
	assert.EqualError(t, err, "cannot delete a permanent admin")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteAdminByIDFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminID := "some-uuid"
	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("Active")

	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)
	mock.ExpectExec("DELETE FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnError(fmt.Errorf("database error"))

	err = s.DeleteAdminByID(adminID)

	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAdminByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminID := "some-uuid"
	adminData := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"}).
		AddRow(adminID, time.Now(), time.Now(), nil, "John Doe", "john@example.com", "Manager", "Active")

	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE id = \\$1").
		WithArgs(adminID).WillReturnRows(adminData)

	admin, err := s.GetAdminByID(adminID)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, admin)
	assert.Equal(t, adminID, admin.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAdminByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	adminID := "some-uuid"

	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE id = \\$1").
		WithArgs(adminID).WillReturnRows(sqlmock.NewRows(nil))

	admin, err := s.GetAdminByID(adminID)

	// assertions
	assert.Error(t, err)
	assert.Nil(t, admin)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestUpdateAdminSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	admin := models.Admin{
		ID:       "some-uuid",
		Name:     "John Doe",
		Email:    "john@example.com",
		Position: "Manager",
		Status:   "Active",
	}

	// Mock successful execution of the update query
	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the function under test
	err = s.UpdateAdmin(admin)

	// Assertions
	assert.NoError(t, err)

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestUpdateAdminNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	admin := models.Admin{
		ID:       "non-existing-uuid",
		Name:     "Jane Doe",
		Email:    "jane@example.com",
		Position: "Developer",
		Status:   "Inactive",
	}

	// Mock the update query with no rows affected
	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// Call the function under test
	err = s.UpdateAdmin(admin)

	// Assertions
	assert.NoError(t, err) // Note: The function itself does not return an error if no rows are affected

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestUpdateAdminQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	admin := models.Admin{
		ID:       "some-uuid",
		Name:     "John Doe",
		Email:    "john@example.com",
		Position: "Manager",
		Status:   "Active",
	}

	// Simulate a database error during the update
	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
		WillReturnError(fmt.Errorf("database error"))

	// Call the function under test
	err = s.UpdateAdmin(admin)

	// Assertions
	assert.Error(t, err)

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
