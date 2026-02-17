package routes

import (
	"yuedsen-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, poseController *controllers.PoseController) *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{

		userGroup.GET("", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("", userController.CreateUser)
		userGroup.POST("/auth/google/callback", userController.GoogleCallback)

	}

	poseGroup := r.Group("/poses")
	{
		poseGroup.GET("", poseController.GetPoses)
		poseGroup.POST("", poseController.CreatePose)
	}


	return r
}
