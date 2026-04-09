package models

import "time"

type BookingSeat struct {
	ID        int       `json:"id"`
	BookingID int       `json:"booking_id"`
	SeatID    string    `json:"seat_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
