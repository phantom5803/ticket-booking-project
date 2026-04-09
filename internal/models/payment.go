package models

import "time"

type Payment struct {
	ID            int     `gorm:"primaryKey"`
	BookingID     int     `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	PaymentMode   string  // UPI, CARD, NETBANKING
	Status        string  // SUCCESS, FAILED, PENDING
	TransactionID string
	CreatedAt     time.Time
}
