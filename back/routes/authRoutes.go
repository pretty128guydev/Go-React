package routes

import (
	"back/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	Controller := controllers.AuthController{}
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", Controller.Register)
		auth.POST("/signin", Controller.Login)
		auth.POST("/find_users", Controller.FindUser)
	}
}
