package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/middleware"
)

func RegisterRoutes(s *gin.Engine) {
	s.GET("/events", getEvents)
	s.GET("/events/:id", singleEvent)
	authRoutes := s.Group("/")
	authRoutes.Use(middleware.Authenticate)
	authRoutes.POST("/events", createEvent)
	authRoutes.PUT("/events/:id", updateEvent)
	authRoutes.DELETE("/events/:id", deleteEvent)
	authRoutes.POST("/events/:id/book", bookEvent)
	authRoutes.GET("/events/bookings", getBookings)
	authRoutes.DELETE("/events/:id/book", cancelEventBooking)
	authRoutes.GET("/user/details",GetUserDetails)
	s.POST("/users/signup", signUp)
	s.POST("/users/login", login)
}


