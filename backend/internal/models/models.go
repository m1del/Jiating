package models

import (
	"time"
)

type Admin struct {
	ID        uint       `json:"id"`                   // primary key
	CreatedAt time.Time  `json:"created_at"`           // time of creation
	UpdatedAt time.Time  `json:"updated_at"`           // time of last update
	DeletedAt *time.Time `json:"deleted_at,omitempty"` // soft delete
	Name      string     `json:"name"`                 // name of the admin
	Email     string     `json:"email"`                // gmail of admin
	Position  string     `json:"position"`             // position is the position of the admin, with Jiating email being primary
	Status    string     `json:"status"`               // active, inactive, hiatus
	Events    []Event    `json:"-"`                    // events is a list of events associated with the admin.
}

type Event struct {
	ID          uint      `json:"id"`           // primary key
	CreatedAt   time.Time `json:"created_at"`   // time of creation
	UpdatedAt   time.Time `json:"updated_at"`   // time of last update
	DeletedAt   time.Time `json:"deleted_at"`   // time of deletion, permenant !!
	AdminID     uint      `json:"admin_id"`     // foreign key to the admin
	EventName   string    `json:"event_name"`   // name of the event
	Date        time.Time `json:"date"`         // date of the event
	Content     string    `json:"content"`      // for text content, ig/yt embed links
	Draft       bool      `json:"draft"`        // if the event is a draft
	PublishedAt time.Time `json:"published_at"` // time of publication, nil if draft
	Admin       Admin     `json:"-"`            // one to one relationship
	Images      []Image   `json:"images"`       // one to many relationship
}

type Image struct {
	ID        uint      `json:"id"`         // primary key
	CreatedAt time.Time `json:"created_at"` // time of creation
	UpdatedAt time.Time `json:"updated_at"` // time of last update
	EventID   uint      `json:"event_id"`   // foreign key to the event
	ImageURL  string    `json:"image_url"`  // url of the image (s3)
}
