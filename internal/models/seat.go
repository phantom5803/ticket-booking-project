package models

type Seat struct {
	ID         int     `json:"id"`
	ShowID     int     `json:"show_id"`
	SeatNumber string  `json:"seat_number"`
	SeatType   string  `json:"seat_type"`
	Price      float64 `json:"price"`
}
