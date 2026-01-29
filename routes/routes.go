package routes

import (
	"yuedsen-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.GET("", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("", userController.CreateUser)
	}

	return r
}
