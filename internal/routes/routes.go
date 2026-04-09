package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phantom5803/ticket-booking-project/internal/middleware"


	"github.com/phantom5803/ticket-booking-project/internal/controllers"
)

func SetupRoutes(r *gin.Engine) {

    // Health check
    r.GET("/health", controllers.HealthCheck)

    // Auth routes
    auth := r.Group("/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{

    // Event routes
    events := protected.Group("/events")
    {
        events.GET("/", controllers.GetEvents)
        events.GET("/:id", controllers.GetEventByID)
    }

    // Show routes
    shows := protected.Group("/shows")
    {
        shows.GET("/:id/seats", controllers.GetSeatsByShow)
    }

    // Seat routes
    seats := protected.Group("/seats")
    {
        seats.POST("/book", controllers.BookSeats)
    }

    // Booking routes
    bookings := protected.Group("/bookings")
    {
        bookings.POST("/", controllers.CreateBooking)
        bookings.GET("/:id", controllers.GetBookingByID)
		bookings.GET("/all",controllers.GetAllBookings)
    }
}
}