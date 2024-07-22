package routes

import "github.com/gin-gonic/gin"

// responsible for registering all routes
// We make RegisterRoutes public to be called from main but the actual handlers are kept in the "routes" package.

func RegisterRoutes(server *gin.Engine) {
	// all get requests for /events will be handled by a handler function. this handler function is automatically
	// given a gin context. We are using a pointer to the single server instance of gin.Engine so there is no need to return anything.

	server.GET("/events", getEvents)
	server.GET("/events/:eventID", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:eventID", updateEvent)
	server.DELETE("/events/:eventID", deleteEvent)
	server.POST("/signup")
}
