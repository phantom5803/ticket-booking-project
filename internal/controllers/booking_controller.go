package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Booking created",
    })
}

func GetBookingByID(c *gin.Context) {
    id := c.Param("id")

    c.JSON(http.StatusOK, gin.H{
        "booking_id": id,
    })
}

func GetAllBookings(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "bookings": []string{"Booking1", "Booking2"},
    })
}