package models

import (
	"time"
)

type Admin struct {
	ID        string       `json:"id"`                   // primary key
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
	ID          string    `json:"id"`           // primary key
	CreatedAt   *time.Time `json:"created_at"`   // time of creation
	UpdatedAt   *time.Time `json:"updated_at"`   // time of last update
	AdminID     int      `json:"admin_id"`     // foreign key to the admin
	EventName   string    `json:"event_name"`   // name of the event
	Date        string `json:"date"`         // date of the event
	Description string    `json:"description"`  // description of the event
	Content     string    `json:"content"`      // for text content, ig/yt embed links
	IsDraft     bool      `json:"is_draft"`     // if the event is a draft
	PublishedAt *time.Time `json:"published_at"` // time of publication, nil if draft
	ImageID      int   `json:"image_id"`       // one to one (for now)
}

// type EventFormRequest struct {
// 	EventID    string `json:"eventID"`
// 	CreatedAt   string `json:"createdAt"`
// 	UpdatedAt   string `json:"updatedAt"`
// 	AdminID     uint `json:"adminID"`
// 	EventName   string `json:"eventName"`
// 	Date		string `json:"date"`
// 	Description string `json:"description"`
// 	Content	 string `json:"content"`
// 	IsDraft	 bool `json:"isDraft"`
// 	PublishedAt string `json:"publishedAt"`
// 	Admin	   string `json:"admin"`
// 	Image	string `json:"image"`
// }

type Image struct {
	ID        uint      `json:"id"`         // primary key
	CreatedAt time.Time `json:"created_at"` // time of creation
	UpdatedAt time.Time `json:"updated_at"` // time of last update
	ImageURL  string    `json:"image_url"`  // url of the image (s3)
}


