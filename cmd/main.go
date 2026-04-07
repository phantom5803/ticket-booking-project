package main

import (
    "log"

    "github.com/phantom5803/ticket-booking-project/config"
    "github.com/phantom5803/ticket-booking-project/internal/routes"
    "github.com/phantom5803/ticket-booking-project/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    // Load env
    config.LoadConfig()

    // Connect DB
    database.Connect()

    // Start server
    r := gin.Default()

    routes.SetupRoutes(r)

    port := config.GetEnv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("🚀 Server running on port", port)
    r.Run(":" + port)
}