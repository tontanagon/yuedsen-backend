package main

import (
	"yuedsen-backend/controllers"
	"yuedsen-backend/repositories"
	"yuedsen-backend/routes"
	"yuedsen-backend/services"
)

func main() {
	// Initialize Repository
	userRepo := repositories.NewUserRepository()

	// Initialize Service
	userService := services.NewUserService(userRepo)

	// Initialize Controller
	userController := controllers.NewUserController(userService)

	// Initialize Router
	r := routes.SetupRouter(userController)

	// Run Server
	r.Run(":8080")
}
