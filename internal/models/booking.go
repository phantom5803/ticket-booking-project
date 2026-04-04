package models

import "time"

type Booking struct {
	ID          string    `json:"id" bson:"_id"`
	ShowID      string    `json:"show_id" bson:"show_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	TotalAmount float64   `json:"total_amount" bson:"total_amount"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	User User ` bson:"foreignkey:UserID"`
	Show Show ` bson:"foreignkey:ShowID"`

}
