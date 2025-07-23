package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AjayKumar-j-s/RecipeSharing/db"
	"github.com/AjayKumar-j-s/RecipeSharing/models"
	"github.com/AjayKumar-j-s/RecipeSharing/routes"
)


func main(){

	server := gin.Default()

	db.InitDB()

	models.CreateTable()

	routes.Registerroutes(server)
	
	server.Run(":8080")

}