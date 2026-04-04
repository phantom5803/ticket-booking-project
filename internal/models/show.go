package models

import "time"

type Show struct {
	ID      string `json:"id" bson:"_id"`
	EventID string `json:"event_id" bson:"event_id"`

	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
	Location  string    `json:"location" bson:"location"`
	Event Event `json:"event" bson:"event"`
}
