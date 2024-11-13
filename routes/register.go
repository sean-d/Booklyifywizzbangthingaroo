package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/models"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	/*
		Grab userId from the context (which ultimately came from the token)
		Grab eventId from the URL parameters

		get specific event based on the id from the URL

		Using the register method for the event model, we register a user for the event.

	*/
	userId := context.GetInt64("userID")

	// the param is a string, so we convert it to an int. we use base 10 since it's decimal and 64 since our
	// event id type in the Event struct is int64.
	eventId, err := strconv.ParseInt(context.Param("eventID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retreieve event by ID."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful."})
}

func cancelRegistration(context *gin.Context) {
	/*
		Grab userId from the context (which ultimately came from the token)
		Grab eventId from the URL parameters

		get specific event based on the id from the URL

		Using the register method for the event model, we register a user for the event.

	*/
	userId := context.GetInt64("userID")

	// the param is a string, so we convert it to an int. we use base 10 since it's decimal and 64 since our
	// event id type in the Event struct is int64.
	eventId, err := strconv.ParseInt(context.Param("eventID"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retreieve event by ID."})
		return
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration from event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Canceled event registration successfully."})
}
