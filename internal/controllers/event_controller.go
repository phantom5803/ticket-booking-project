package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "events": []string{"Event1", "Event2"},
    })
}

func GetEventByID(c *gin.Context) {
    id := c.Param("id")

    c.JSON(http.StatusOK, gin.H{
        "event_id": id,
    })
}