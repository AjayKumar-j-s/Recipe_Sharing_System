package routes

import "github.com/gin-gonic/gin"

func Registerroutes(server *gin.Engine){

	server.POST("/auth/register",RegisterUser)

}