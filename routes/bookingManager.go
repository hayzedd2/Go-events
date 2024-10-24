package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/Go-events/models"
)

func bookEvent(c *gin.Context) {
	userId := c.GetString("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse eventid",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not parse eventid",
		})
		return
	}
	err = event.Book(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not book event",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event booked!"})

}

func cancelEventBooking(c *gin.Context) {
	userId := c.GetString("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse eventid",
		})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not parse eventid",
		})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelBooking(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not cancel event booking",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event booking cancelled!"})
}
