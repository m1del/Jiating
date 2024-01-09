package tests

import (
	"backend/internal/database"
	"backend/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// CreateEvent tests
func TestCreateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := database.New(db)

	// sample event and admin IDs
	event := models.Event{
		EventName:   "Sample Event",
		Date:        "2024-01-10",
		Description: "This is a sample event",
		Content:     "Event content here",
		IsDraft:     false,
		PublishedAt: nil,
		Images:      []models.EventImage{{ID: "image1"}},
	}
	adminIDs := []string{"admin1", "admin2"}

	// set up expectations
	mock.ExpectBegin()

	// mock insert event query
	rows := sqlmock.NewRows([]string{"id"}).AddRow("event1")
	mock.ExpectQuery("INSERT INTO events").
		WithArgs(event.EventName, event.Date, event.Description, event.Content, event.IsDraft, event.PublishedAt).
		WillReturnRows(rows)

	// mock insert admin-event associations
	for _, adminID := range adminIDs {
		mock.ExpectExec("INSERT INTO event_authors").
			WithArgs(adminID, "event1").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	// mock adding image to event
	for _, img := range event.Images {
		mock.ExpectExec("INSERT INTO images").
			WithArgs(img.ImageURL, sqlmock.AnyArg(), "event1").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	// mock reset isDisplay for all images of this event
	mock.ExpectExec("UPDATE images SET is_display = false WHERE event_id =").
		WithArgs("event1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Mock setting the first image as display image
	if len(event.Images) > 0 {
		// set isDisplay to true for the chosen image
		mock.ExpectExec("UPDATE images SET is_display = true WHERE id =").
			WithArgs(event.Images[0].ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	mock.ExpectCommit()

	_, err = s.CreateEvent(event, adminIDs)

	// Assert there was no error and all expectations were met
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// UpdateEventByID tests
func TestUpdateDynamicEventFields(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := database.New(db)

	// expectations
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE events SET").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectRollback()

	// mock transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	eventID := "event1"
	updatedData := map[string]interface{}{
		"event_name":  "Updated Event Name",
		"description": "Updated Description",
	}

	err = s.UpdateDynamicEventFields(tx, eventID, updatedData)
	assert.NoError(t, err)

	// rollback
	err = tx.Rollback()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEventAuthorship(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := database.New(db)

	// epectations
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT EXISTS").WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectExec("INSERT INTO event_authors").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectRollback()

	// mock transaction
	tx, err := db.Begin()
	assert.NoError(t, err)

	eventID := "event1"
	editorAdminID := "admin1"

	err = s.UpdateEventAuthorship(tx, eventID, editorAdminID)
	assert.NoError(t, err)

	// rollback
	err = tx.Rollback()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateEventByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := database.New(db)

	// sample UpdateEventRequest and event ID
	eventID := "event1"
	req := models.UpdateEventRequest{
		UpdatedData:     map[string]interface{}{"event_name": "Updated Event", "description": "Updated Description"},
		NewImages:       []models.EventImage{{ID: "newImage1"}},
		RemovedImageIDs: []string{"oldImage1"},
		NewDisplayImage: "newImage1",
		EditorAdminID:   "admin1",
	}

	mock.ExpectBegin()

	// dynamic update query for the event
	mock.ExpectExec("UPDATE events SET").
		WithArgs("Updated Event", "Updated Description", sqlmock.AnyArg(), eventID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// addition of new images
	for _, img := range req.NewImages {
		mock.ExpectExec("INSERT INTO images").
			WithArgs(img.ImageURL, img.IsDisplay, eventID).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	// removal of old images
	for _, imgID := range req.RemovedImageIDs {
		mock.ExpectExec("DELETE FROM images WHERE").
			WithArgs(imgID).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	// reset isDisplay for all images of this event
	mock.ExpectExec("UPDATE images SET is_display = false WHERE event_id =").
		WithArgs(eventID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// setting new display image
	if req.NewDisplayImage != "" {
		// set isDisplay to true for the chosen image
		mock.ExpectExec("UPDATE images SET is_display = true WHERE id =").
			WithArgs(req.NewDisplayImage).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}

	// check and update of event authorship
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(req.EditorAdminID, eventID).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
	mock.ExpectExec("INSERT INTO event_authors").
		WithArgs(req.EditorAdminID, eventID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = s.UpdateEventByID(eventID, req)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
