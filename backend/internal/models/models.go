package models

import (
	"time"
)

type Admin struct {
	ID        string        `json:"id"`                   // primary key, UUID
	CreatedAt time.Time     `json:"created_at"`           // time of creation
	UpdatedAt time.Time     `json:"updated_at"`           // time of last update
	DeletedAt *time.Time    `json:"deleted_at,omitempty"` // soft delete
	Name      string        `json:"name"`                 // name of the admin
	Email     string        `json:"email"`                // gmail of admin
	Position  string        `json:"position"`             // position is the position of the admin, with Jiating email being primary
	Status    string        `json:"status"`               // active, inactive, hiatus
	Events    []EventAuthor `json:"events"`               // events is a list of events associated with the admin.
}

type Event struct {
	ID          string       `json:"id"`           // primary key, UUID
	CreatedAt   time.Time    `json:"created_at"`   // time of creation
	UpdatedAt   time.Time    `json:"updated_at"`   // time of last update
	EventName   string       `json:"event_name"`   // name of the event
	Date        string       `json:"date"`         // date of the event
	Description string       `json:"description"`  // description of the event
	Content     string       `json:"content"`      // for text content, ig/yt embed links
	IsDraft     bool         `json:"is_draft"`     // if the event is a draft
	PublishedAt *time.Time   `json:"published_at"` // time of publication, nil if draft
	Images      []EventImage `json:"images"`       // images associated with the event
	Authors     []Admin      `json:"authors"`      // authors associated with the event
}

type EventAuthor struct {
	AdminID string `json:"admin_id"` // foreign key to admin
	EventID string `json:"event_id"` // foreign key to event
}

type EventImage struct {
	ID        uint      `json:"id"`         // primary key
	CreatedAt time.Time `json:"created_at"` // time of creation
	UpdatedAt time.Time `json:"updated_at"` // time of last update
	ImageURL  string    `json:"image_url"`  // url of the image (s3)
	IsDisplay bool      `json:"is_display"` // if the image is the display image for the event
}
