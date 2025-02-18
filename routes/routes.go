package routes

import (
	"go/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.Engine){

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events",createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent )
	authenticated.DELETE("/events/:id/register",cancelRegistration )

	server.GET("/events", getEvents)
	server.GET("/events/:id",getEvent )
	server.POST("/signup", signUp)
	server.POST("/login", login)

}