package tests

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

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

	// mocking email existence check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
		WithArgs(admin.Email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// mocking insert query
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

	// mocking email existence check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
		WithArgs(admin.Email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// set up the mock to expect a QueryRow and return an error
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

func TestCreateAdminDuplicateEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := database.New(db)

	duplicateEmail := "jiating.lion.dragon@gmail.com"
	admin := models.Admin{
		Name:     "doop",
		Email:    duplicateEmail,
		Position: "Founder",
		Status:   "permanent",
	}

	// mocking email existence check, return true for duplicate email
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
		WithArgs(admin.Email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// mocking insert query
	id, err := s.CreateAdmin(admin)

	assert.Error(t, err)
	assert.Empty(t, id, "Expected ID to be empty when an error occurs")
	assert.Contains(t, err.Error(), "email already exists", "Expected error message for duplicate email")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCreateAdminInvalidEmail(t *testing.T) {
	s := database.New(nil)

	admin := models.Admin{
		Name:     "invalid email",
		Email:    "invalid-email",
		Position: "invalid email",
		Status:   "active",
	}
	id, err := s.CreateAdmin(admin)
	assert.Error(t, err)
	assert.Empty(t, id, "Expected ID to be empty when an error occurs")
	assert.Contains(t, err.Error(), "invalid email", "Expected error message for invalid email")
}

// simulating a unique violation error -> race condition where the email becomes duplicated between the check and the insert
func TestCreateAdminUniqueViolation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := database.New(db)

	admin := models.Admin{
		Name:     "unique violation",
		Email:    "john@mail.com",
		Position: "unique violation",
		Status:   "active",
	}

	// mock database error during email existence check
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
		WithArgs(admin.Email).
		WillReturnError(fmt.Errorf("database error"))

	id, err := s.CreateAdmin(admin)

	assert.Error(t, err)
	assert.Empty(t, id, "Expected ID to be empty when an error occurs")
	assert.Contains(t, err.Error(), "database error", "Expected error message for database error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

// if two requestse trying to create an admin with the same email
func TestCreateAdminRaceCondition(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error %s was not expected when opening a stub database connection", err)
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

	// mocking email existence check
	// assuming that the first call will be successful and subsequent calls will fail
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
		WithArgs(admin.Email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// set up mock for successful insert
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO admins")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), admin.Name, admin.Email, admin.Position, admin.Status).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("some-uuid"))

	var wg sync.WaitGroup
	successCount := 0
	failureCount := 0
	for i := 0; i < 5; i++ { // simulate 5 concurrent requests
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.CreateAdmin(admin)
			if err != nil {
				failureCount++
			} else {
				successCount++
			}
		}()
	}
	wg.Wait()
	// check only one insert was successful
	assert.Equal(t, 1, successCount, "Only one admin creation should be successful")
	assert.Equal(t, 4, failureCount, "Nine admin creations should fail")
}

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
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE deleted_at IS NULL")).WillReturnRows(rows)
	resultAdmins, err := s.GetAllAdmins()

	assert.NoError(t, err)
	assert.NotNil(t, resultAdmins)
	assert.Equal(t, len(admins), len(resultAdmins), "Expected number of admins does not match")
	for i, admin := range resultAdmins {
		assert.Equal(t, admins[i], admin, "Admin data does not match")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetAllAdminsDatabaseQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error %s was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := database.New(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE deleted_at IS NULL")).
		WillReturnError(fmt.Errorf("query error"))
	// Simulte a database error during querying
	admins, err := s.GetAllAdmins()

	assert.Error(t, err)
	assert.Nil(t, admins)
	assert.Contains(t, err.Error(), "query error", "Expected error message for database error")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

// func TestGetAdminByEmail(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	// Test data
// 	testEmail := "test@example.com"
// 	testAdmin := models.Admin{
// 		ID:        "test-uuid",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 		Name:      "Test Admin",
// 		Email:     testEmail,
// 		Position:  "Test Position",
// 		Status:    "Active",
// 	}

// 	// mocking email existence chec
// 	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)")).
// 		WithArgs(testAdmin.Email).
// 		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

// 	// Setup mock expectation
// 	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"}).
// 		AddRow(testAdmin.ID, testAdmin.CreatedAt, testAdmin.UpdatedAt, testAdmin.DeletedAt, testAdmin.Name, testAdmin.Email, testAdmin.Position, testAdmin.Status)

// 	mock.ExpectQuery("^SELECT (.+) FROM admins WHERE email = \\$1$").
// 		WithArgs(testEmail).
// 		WillReturnRows(rows)

// 	admin, err := s.GetAdminByEmail(testEmail)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, admin)
// 	assert.Equal(t, testAdmin, *admin)

// 	// check if all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetAdminByEmailQueryError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminEmail := "error@example.com"

// 	// Simulate a database error during querying
// 	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE email = \\$1").
// 		WithArgs(adminEmail).WillReturnError(fmt.Errorf("database error"))

// 	// Call the function under test
// 	admin, err := s.GetAdminByEmail(adminEmail)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, admin)

// 	// Check if all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestDeleteAdminByIDSuccess(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminID := "some-uuid"
// 	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("Active")

// 	// mock query to check admin status
// 	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)

// 	// mock successful execution of the delete query
// 	mock.ExpectExec("DELETE FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnResult(sqlmock.NewResult(1, 1))

// 	// call the function under test
// 	err = s.DeleteAdminByID(adminID)

// 	assert.NoError(t, err)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestDeleteAdminByIDPermanent(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminID := "some-uuid"
// 	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("permanent")

// 	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)
// 	err = s.DeleteAdminByID(adminID)

// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "cannot delete a permanent admin")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestDeleteAdminByIDFailure(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminID := "some-uuid"
// 	statusRows := sqlmock.NewRows([]string{"status"}).AddRow("Active")

// 	mock.ExpectQuery("SELECT status FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnRows(statusRows)
// 	mock.ExpectExec("DELETE FROM admins WHERE id = \\$1").WithArgs(adminID).WillReturnError(fmt.Errorf("database error"))

// 	err = s.DeleteAdminByID(adminID)

// 	assert.Error(t, err)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetAdminByIDSuccess(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminID := "some-uuid"
// 	adminData := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "email", "position", "status"}).
// 		AddRow(adminID, time.Now(), time.Now(), nil, "John Doe", "john@example.com", "Manager", "Active")

// 	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE id = \\$1").
// 		WithArgs(adminID).WillReturnRows(adminData)

// 	admin, err := s.GetAdminByID(adminID)

// 	// assertions
// 	assert.NoError(t, err)
// 	assert.NotNil(t, admin)
// 	assert.Equal(t, adminID, admin.ID)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetAdminByIDNotFound(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	adminID := "some-uuid"

// 	mock.ExpectQuery("SELECT id, created_at, updated_at, deleted_at, name, email, position, status FROM admins WHERE id = \\$1").
// 		WithArgs(adminID).WillReturnRows(sqlmock.NewRows(nil))

// 	admin, err := s.GetAdminByID(adminID)

// 	// assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, admin)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestUpdateAdminSuccess(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	admin := models.Admin{
// 		ID:       "some-uuid",
// 		Name:     "John Doe",
// 		Email:    "john@example.com",
// 		Position: "Manager",
// 		Status:   "Active",
// 	}

// 	// Mock successful execution of the update query
// 	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
// 		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	// Call the function under test
// 	err = s.UpdateAdmin(admin)

// 	// Assertions
// 	assert.NoError(t, err)

// 	// Check if all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestUpdateAdminNotFound(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	admin := models.Admin{
// 		ID:       "non-existing-uuid",
// 		Name:     "Jane Doe",
// 		Email:    "jane@example.com",
// 		Position: "Developer",
// 		Status:   "Inactive",
// 	}

// 	// Mock the update query with no rows affected
// 	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
// 		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
// 		WillReturnResult(sqlmock.NewResult(0, 0))

// 	// Call the function under test
// 	err = s.UpdateAdmin(admin)

// 	// Assertions
// 	assert.NoError(t, err) // Note: The function itself does not return an error if no rows are affected

// 	// Check if all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestUpdateAdminQueryError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	s := database.New(db)

// 	admin := models.Admin{
// 		ID:       "some-uuid",
// 		Name:     "John Doe",
// 		Email:    "john@example.com",
// 		Position: "Manager",
// 		Status:   "Active",
// 	}

// 	// Simulate a database error during the update
// 	mock.ExpectExec("UPDATE admins SET name = \\$1, email = \\$2, position = \\$3, status = \\$4, updated_at = \\$5 WHERE id = \\$6").
// 		WithArgs(admin.Name, admin.Email, admin.Position, admin.Status, sqlmock.AnyArg(), admin.ID).
// 		WillReturnError(fmt.Errorf("database error"))

// 	// Call the function under test
// 	err = s.UpdateAdmin(admin)

// 	// Assertions
// 	assert.Error(t, err)

// 	// Check if all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }
