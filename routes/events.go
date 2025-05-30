package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/models"
	"net/http"
	"strconv"
)

// routes/event.go: handles all request handler functions that deal with events

func getEvent(context *gin.Context) {
	// the param is a string, so we convert it to an int. we use base 10 since it's decimal and 64 since our
	// event id type in the Event struct is int64.
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	event, err := models.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find event"})
	}

	context.JSON(http.StatusOK, event)

}
func getEvents(context *gin.Context) {
	// get all events

	// to return json, one simply does not return json. one must invoke JSON method on the gin context struct that
	// was passed in and provide a status code and data. tasty tasty data.
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get all events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	/*
		Using ShouldBindJSON to create a json object based on the struct def of events.Event.
		If there's an error, a json representation of an error is returned

		we grab the userId from the context as it was set in the Authenticate middleware.
		it is a float64 so we use context.GetFLoat64() to retrieve the value for the supplied key.


		if we are not able to create an event, we return an error and return out

	*/

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not create event."})
		return
	}

	userId := context.GetInt64("userID")
	event.UserID = userId

	// if everything works...
	err = event.Save()
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	// the param is a string, so we convert it to an int. we use base 10 since it's decimal and 64 since our
	// event id type in the Event struct is int64.
	eventId, err := strconv.ParseInt(context.Param("eventID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userID")
	event, err := models.GetEvent(eventId) // we are not using the resulting event, only checking that it was queried successfully. will be using it later on. leaving _ until needed

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event // new event with updated information to be used to update the row in db

	err = context.ShouldBindJSON(&updatedEvent) // creating json object
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not update event"})
	}

	updatedEvent.ID = eventId // the eventID we received as a param to the updateEvent function
	updatedEvent.UserID = userId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	// the param is a string, so we convert it to an int. we use base 10 since it's decimal and 64 since our
	// event id type in the Event struct is int64.
	eventID, err := strconv.ParseInt(context.Param("eventID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userID")
	event, err := models.GetEvent(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find event"})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.DeleteEvent()

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted", "event": event})

}
