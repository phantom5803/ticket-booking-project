package models


type Seat struct {
	ID         string    `json:"id" bson:"_id"`
	ShowID     string    `json:"show_id" bson:"show_id"`
	SeatNumber string    `json:"seat_number" bson:"seat_number"`
	SeatType   string    `json:"seat_type" bson:"seat_type"`
	Price      float64   `json:"price" bson:"price"`
	
}
