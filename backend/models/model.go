package models

import (
	"time"

	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

// Admin represents a Jaiting admin. Jiating email is perma linked to the admin, and users are added based on board membership.
type Admin struct {
	gorm.Model           // adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	AdminID      uint    `gorm:"primaryKey;autoIncrement" json:"admin_id"`   // primary key
	Name         string  `gorm:"type:varchar(100)" json:"name"`              // name of the admin
	Email        string  `gorm:"type:varchar(100);uniqueIndex" json:"email"` // personal email
	PasswordHash string  `gorm:"not null"`                                   // password to login to admin dashboard
	Position     string  `gorm:"type:varchar(100)" json:"position"`          // position in the organization
	Status       string  `gorm:"type:varchar(10)" json:"status"`             // active, inactive, hiatus
	Events       []Event `json:"-"`                                          // events associated with the admin
}

// Event describes a Jiating event. Events are created by admins and are displayed on the website.
type Event struct {
	gorm.Model
	EventID   uint      `gorm:"primaryKey;autoIncrement" json:"event_id"` // primary key
	AdminID   uint      `gorm:"not null" json:"admin_id"`                 // foreign key to the admin
	EventName string    `gorm:"type:varchar(100)" json:"event_name"`      // name of the event
	Date      time.Time `json:"date"`                                     // date of the event
	Content   string    `gorm:"type:text" json:"content"`                 // for text content, ig/yt embed links
	Admin     Admin     `gorm:"foreignKey:AdminID" json:"-"`              // one to one relationship
	Images    []Image   `gorm:"foreignKey:EventID" json:"images"`         // one to many relationship
}

// Events can contain images. Images are uploaded to s3 and the url is stored in the database.
type Image struct {
	gorm.Model
	ImageID  uint   `gorm:"primaryKey;autoIncrement" json:"image_id"` // primary key
	EventID  uint   `gorm:"not null" json:"event_id"`                 // foreign key to the event
	ImageURL string `gorm:"type:varchar(255)" json:"image_url"`       // url of the image (s3)
}
