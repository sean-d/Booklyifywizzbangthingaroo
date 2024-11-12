package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/utils"
	"net/http"
)

func Authenticate(context *gin.Context) {
	/*
		"Authorization" in the header is where things such as tokens will reside.

		Since this sits in the middle of a handler, we need to ensure that any errors that require
		stopping are actually stopping. We use AbortWithStatusJSON to ensure Abort() is called and a JSON response sent.

		if token is blank, we Abort and stop all future handlers from running
		if the token cannot be verified, we Abort...


		context.Set() to ensure we are setting the userId for the context so it can be used in events.

		If we have no errors, we finish by calling context.Next() to ensure the next request handler in line
		will execute correctly.
	*/
	token := context.Request.Header.Get("Authorization")

	// check for no token
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized: empty token"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized: cannot validate"})
		return
	}

	context.Set("userID", userId)
	context.Next()
}
