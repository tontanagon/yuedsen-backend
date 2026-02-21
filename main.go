package main

import (
	"log"
	"yuedsen-backend/controllers"
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
	"yuedsen-backend/routes"
	"yuedsen-backend/services"

	"yuedsen-backend/repositories/dbmanager"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Database Connection
	db := dbmanager.Connect()

	// Auto Migrate
	err = db.AutoMigrate(&models.User{}, &models.PoseCategory{}, &models.Pose{}, &models.Plan{}, &models.UserProcess{}, &models.UserToken{}, &models.PoseLandmark{})
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

	// Initialize Pose Stack
	poseRepo := repositories.NewPoseRepository(db)
	poseService := services.NewPoseService(poseRepo)
	poseController := controllers.NewPoseController(poseService)

	// Initialize Game Stack
	gameRepo := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo)
	gameController := controllers.NewGameController(gameService)

	// Initialize Router
	r := routes.SetupRouter(userController, poseController, gameController)

	// Run Server
	r.Run(":8080")
}
