package tests

// CRUD tests

// func TestAddImageToEvent(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
// 	s := database.New(db)

// 	eventID := "event-uuid"
// 	image := models.EventImage{
// 		ID:        "image-uuid",
// 		CreatedAt: time.Now(),
// 		ImageURL:  "https://image-url.com",
// 		AltText:   "example image",
// 		IsDisplay: true,
// 	}

// 	mock.ExpectBegin() // begin transaction

// 	//mock insert query
// 	mock.ExpectExec(regexp.QuoteMeta(`
// 		INSERT INTO event_images(
// 			id, created_at, image_url,
// 			alt_text, is_display, event_id
// 		) VALUES ($1, $2, $3, $4, $5, $6)`)).WithArgs(
// 		image.ID, sqlmock.AnyArg(), image.ImageURL,
// 		image.AltText, image.IsDisplay, eventID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	tx, err := db.Begin()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when beginning a transaction", err)
// 	}

// 	err = s.AddImageToEventTx(tx, image, eventID)
// 	assert.NoError(t, err)

// 	mock.ExpectCommit() // commit transaction

// 	if err := tx.Commit(); err != nil {
// 		t.Fatalf("an error '%s' was not expected when committing a transaction", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
