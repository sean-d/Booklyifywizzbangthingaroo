package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/middlewares"
)

// responsible for registering all routes
// We make RegisterRoutes public to be called from main but the actual handlers are kept in the "routes" package.

func RegisterRoutes(server *gin.Engine) {
	/*
		all get requests will be handled by a handler function. this handler function is automatically
		given a gin context. We are using a pointer to the single server instance of gin.Engine so there is no need to return anything.

		we want to use the middlewares.authenticate for the routes that require protecting. we create a group for this
		and add the routes that will use this functionality...so we load that first, then the routes
		The group is "/" since that is the common thread for all routes.

	*/

	authenticatedMiddleware := server.Group("/")
	authenticatedMiddleware.Use(middlewares.Authenticate)
	authenticatedMiddleware.POST("/events", createEvent)
	authenticatedMiddleware.PUT("/events/:eventID", updateEvent)
	authenticatedMiddleware.DELETE("/events/:eventID", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:eventID", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
