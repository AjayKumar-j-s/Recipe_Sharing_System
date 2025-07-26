package routes

import (
	"net/http"
	"strconv"

	"github.com/AjayKumar-j-s/RecipeSharing/models"
	"github.com/gin-gonic/gin"
)

func Addrecipe(context *gin.Context) {

	var rep models.Recipe
	
	err := context.ShouldBindJSON(&rep)
	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldn't Parse"})
		return
	}

	id := context.GetInt64("uid")
	
	rep.UserID = id

	err = rep.Save()

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldn't Parse"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Recipe Created Successful"})



}



func Getrecipes(contex *gin.Context){
	
	rep,err := models.GetAllrecipes()

	if(err != nil){
		contex.JSON(http.StatusInternalServerError,gin.H{"message":"Unable to get recipes"})
		return
	}

	contex.JSON(http.StatusOK,gin.H{"recipes":rep})

}

func GetrecipeById(context *gin.Context){
	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Invalid request"})
		return
	}

	res,err := models.GetById(id)

	if( err !=nil ){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Couldn't fetch try again later"})
		return
	}
	context.JSON(http.StatusOK,gin.H{"recipe":res})




}


func UpdateRecipe(context *gin.Context){

	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"couid not fetch..Try Again"})
		return
	}

	uid := context.GetInt64("uid")
	recp,err := models.GetById(id)

	print(uid,recp.UserID)
	if(recp.UserID != uid){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized to Update Recipe"})
		return
	}

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch Recipe"})
		return
	}
	


	var r models.Recipe

	err = context.ShouldBindJSON(&r)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not fetch data from req"})
		return
	}

	r.ID = id

	err = r.UpdateRec()

	if(err != nil){

		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch Recipe"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Updated Successfully"})
}



func DeleteRecipe(context *gin.Context){

	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"couid not fetch..Try Again"})
		return
	}

	uid := context.GetInt64("uid")
	recp,err := models.GetById(id)

	print(uid,recp.UserID)
	if(recp.UserID != uid){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized to Update Recipe"})
		return
	}

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch Recipe"})
		return
	}
	


	var r models.Recipe

	err = context.ShouldBindJSON(&r)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not fetch data from req"})
		return
	}

	r.ID = id

	err = r.DeleteRec()

	if(err != nil){

		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch Recipe"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Deleted Successfully"})
	

}