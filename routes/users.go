package routes

import (
	"go/rest-api/models"
	"go/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


	func signUp(context *gin.Context){

		var user models.User
		err := context.ShouldBindJSON(&user)
		if err!=nil {
			context.JSON(http.StatusBadRequest , gin.H{"message":"Could not parse event ID"})
			return
		}
		err = user.Save()
		if err!=nil {
			context.JSON(http.StatusBadRequest , gin.H{"message":"Could not save user "})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message":"User created sucessfully"})


	}

	func login(context *gin.Context){
		var user models.User
		err := context.ShouldBindJSON(&user)
		if err!=nil {
			context.JSON(http.StatusBadRequest , gin.H{"message":"Could not parse event ID"})
			return
		}
		err = user.AuthenticateUser()

		if err!=nil{
			context.JSON(http.StatusUnauthorized , gin.H{"message":"Could not authenticate user"})
		return
		}

		token , err := utils.GenerateToken(user.Email, user.ID)

		if err!=nil{
            context.JSON(http.StatusInternalServerError , gin.H{"message":"Could not generate token"})
            return
        }

		context.JSON(http.StatusOK , gin.H{"message":"User logged in sucessfully", "token":token})


	}