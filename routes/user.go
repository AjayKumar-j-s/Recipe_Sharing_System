package routes

import (
	"net/http"

	"github.com/AjayKumar-j-s/RecipeSharing/models"
	"github.com/AjayKumar-j-s/RecipeSharing/utils"
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


func Login(context *gin.Context){

	var u models.User

	err := context.ShouldBindJSON(&u)

	if(err !=nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not login...Try Again"})
		return
	}

	err = u.ValidatePAssword()


	if(err !=nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Incorrect Password...Try Again"})
		return
	}


	token,err := utils.GenerateToken(u.Email,u.UserID)




	
	if(err !=nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not authenticate"})
		return 
	}

	context.JSON(200,gin.H{"message":"Login Successful","token":token})






}