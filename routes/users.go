package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/models"
	"github.com/sean-d/Booklyifywizzbangthingaroo/utils"
	"net/http"
)

func login(context *gin.Context) {
	/*
		Take json, bind to a user struct
		take user and validate the password associated with email is correct
			if not, return unauth status code via json and the error message

		successful auth will return a login success message and the created token.
	*/

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse requested data."})
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}

func signup(context *gin.Context) {
	/*
		Create a user, bind that to the context, and take the json fields passed in
		and save them as a new user in the user table
	*/
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse requested data.."})
		return // we return to prevent the rest of this from proceeding.
	}

	// if everything works...
	err = user.Save()

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save/create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": user})
}
