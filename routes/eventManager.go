package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/hayzedd2/eventease-be/models"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not  fetch events",
		})
		return
	}
	c.JSON(http.StatusOK, events)
}

func singleEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse data!",
		})
		return
	}
	userId := c.GetString("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "event created!",
		"event":   event,
	})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}
	userId := c.GetString("userId")
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update event!",
		})
		return
	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse request data",
		})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated!",
	})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	userId := c.GetString("userId")
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to delete event!",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}
	err = event.DELETE()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})

}
