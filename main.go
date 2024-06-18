package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/models"
	"net/http"
)

func main() {
	server := gin.Default()

	// all get requests for /events will be handled by a handler function. this handler function is automatically
	// given a gin context
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	err := server.Run(":8080")

	if err != nil {
		panic(err)
	}
}

func getEvents(context *gin.Context) {
	// to return json, one simply does not return json. one must invoke JSON method on the gin context struct that
	// was passed in and provide a status code and data. tasty tasty data.
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	/*
		Using ShouldBindJSON to create a json object based on the
		struct def of events.Event. If there's an error, a json representation of an error is returned

	*/
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return // we return to prevent the rest of this from proceeding.
	}

	// dummy entries for now until we move to a proper db
	event.ID = 1
	event.UserID = 1

	// if everything works...
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
