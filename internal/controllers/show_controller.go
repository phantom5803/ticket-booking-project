package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetSeatsByShow(c *gin.Context) {
    showID := c.Param("id")

    c.JSON(http.StatusOK, gin.H{
        "show_id": showID,
        "seats":   []string{"A1", "A2", "A3"},
    })
}