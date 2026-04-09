package models

import "time"

type Booking struct {
	ID          int       `json:"id"`
	ShowID      int       `json:"show_id"`
	UserID      int       `json:"user_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	User        User
	Show        Show
}
