package routes

import (
	"net/http"

	"github.com/AjayKumar-j-s/RecipeSharing/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context){

	var user models.User

	err := context.ShouldBindJSON(&user)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not signup try again"})
		return
	}

	err = user.Save()

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not register"})
		return
	}


	context.JSON(http.StatusOK,gin.H{"message":"Signedup  Successful"})



}