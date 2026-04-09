package models

import "time"

type Show struct {
	ID        int       `json:"id" gorm:"primaryKey;type:integer"`
	EventID   int       `json:"event_id" gorm:"type:integer"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Location  string    `json:"location"`
	Event     Event     `json:"event"`
}
