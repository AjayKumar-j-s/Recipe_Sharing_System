package routes

import (
	"github.com/AjayKumar-j-s/RecipeSharing/middleware"
	"github.com/gin-gonic/gin"
)

func Registerroutes(server *gin.Engine){

	server.POST("/auth/register",RegisterUser)
	server.POST("/login",Login)

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/recipe",Addrecipe)
	authenticate.GET("/recipes",Getrecipes)
	authenticate.GET("/recipe/:id",GetrecipeById)
	authenticate.PUT("/recipe/:id",UpdateRecipe)
	authenticate.DELETE("/recipe/:id",DeleteRecipe)




}