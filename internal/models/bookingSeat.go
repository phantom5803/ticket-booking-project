package models

import "time"

type BookingSeat struct {
	ID        string    `json:"id" bson:"_id"`
	BookingID string    `json:"booking_id" bson:"booking_id"`
	SeatID    string    `json:"seat_id" bson:"seat_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
