package database

import (
    "fmt"
    "log"

    "github.com/phantom5803/ticket-booking-project/config"
	"github.com/phantom5803/ticket-booking-project/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        config.GetEnv("DB_HOST"),
        config.GetEnv("DB_USER"),
        config.GetEnv("DB_PASSWORD"),
        config.GetEnv("DB_NAME"),
        config.GetEnv("DB_PORT"),
        config.GetEnv("DB_SSLMODE"),
    )
	DB.AutoMigrate(
        &models.User{},
        &models.Event{},
        &models.Show{},
        &models.Seat{},
        &models.Booking{},
        &models.BookingSeat{},
        &models.Payment{},
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("✅ Database connected")
}