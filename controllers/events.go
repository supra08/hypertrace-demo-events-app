package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Event created successfully!"})
}

func FetchAllEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "events bla bla"})
}

func FetchEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "event bla bla"})
}

func DeleteEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Event deleted successfully!"})
}

func UpdateEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Event updated successfully!"})
}
