package routes

import (
	"fmt"
	"go/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func getEvents(context *gin.Context){
	events , err := models.GetAllEvents()

	fmt.Print(err)
	if err!=nil {
	   context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not parse request data"})
	     return 
	 }

	context.JSON(http.StatusOK, events) 
}

func createEvent(context *gin.Context){
	var newEvent models.Event
    err := context.ShouldBindJSON(&newEvent)
	
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
	return
	}

	userId := context.GetInt64("userId")
	newEvent.UserID = userId

    err = newEvent.Save()

	if err!=nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not create event "})
		return
		 }

	 context.JSON(http.StatusOK, newEvent)
   
}


func getEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not parse event ID"})
		return
	}
	event , err := models.GetEventByID(eventId)

	if err!=nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not parse event ID"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err!=nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not fetch event"})
		return
	}

	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized , gin.H{"message":"Unauthorized to get event"})
		return
	}



	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err!=nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not update"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err!=nil {
        context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not update event"})
        return
    }
	context.JSON(http.StatusOK , gin.H{"message":"Request updated sucessfully"})

}

func deleteEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("id"),10,64)

	if err!=nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not parse event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event , err := models.GetEventByID(eventId)

	if err!=nil {
		context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not fetch event"})
		return
	}
	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized , gin.H{"message":"Unauthorized to delete event"})
		return
	}

	err = event.Delete()

	if err!=nil {
		context.JSON(http.StatusBadRequest , gin.H{"message":"Could not delete event "})
		return
	}

	context.JSON(http.StatusOK , gin.H{"message":"Event deleted sucessfully"})

}