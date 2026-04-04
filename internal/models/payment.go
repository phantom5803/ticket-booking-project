package models

import "time"

type Payment struct {
    ID            uint      `gorm:"primaryKey"`
    BookingID     uint      `gorm:"not null"`
    Amount        float64   `gorm:"not null"`
    PaymentMode   string    // UPI, CARD, NETBANKING
    Status        string    // SUCCESS, FAILED, PENDING
    TransactionID string
    CreatedAt     time.Time
}