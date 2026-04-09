package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func BookSeats(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Seats booked (mock)",
    })
}