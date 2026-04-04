package models

import "time"

type Event struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	StartDate   time.Time `json:"start_date" bson:"start_date"`
	EndDate     time.Time `json:"end_date" bson:"end_date"`
	Location    string    `json:"location" bson:"location"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
