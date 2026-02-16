package main

import (
	"log"
	"yuedsen-backend/controllers"
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
	"yuedsen-backend/routes"
	"yuedsen-backend/services"

	"yuedsen-backend/repositories/dbmanager"
)

func main() {
	// Database Connection
	db := dbmanager.Connect()

	// Auto Migrate
	err := db.AutoMigrate(&models.User{}, &models.PoseCategory{}, &models.Pose{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")

	// Initialize Repository
	userRepo := repositories.NewUserRepository(db)

	// Initialize Service
	userService := services.NewUserService(userRepo)

	// Initialize Controller
	userController := controllers.NewUserController(userService)

	// Initialize Router
	r := routes.SetupRouter(userController)

	// Run Server
	r.Run(":8080")
}
