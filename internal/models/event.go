package models

import "time"

type Event struct {
	ID          int    `gorm:"primaryKey;type:integer"`
	Title       string `gorm:"not null"`
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
