package routes

import (
	"net/http"
	"os"
	"strings"
	"yuedsen-backend/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// JWTMiddleware — ตรวจ Bearer token และ inject user_id เข้า context
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "fallback_secret_for_dev"
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// inject claims เข้า gin context ให้ handler ใช้ได้
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
			c.Set("name", claims["name"])
		}

		c.Next()
	}
}

func SetupRouter(userController *controllers.UserController, poseController *controllers.PoseController, gameController *controllers.GameController) *gin.Engine {
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// === Public routes (ไม่ต้อง token) ===
	userGroup := r.Group("/users")
	{
		userGroup.POST("", userController.CreateUser)               // upsert user
		userGroup.POST("/token", userController.GetToken)           // get JWT by email
		userGroup.POST("/auth/google/callback", userController.GoogleCallback)
	}

	// === Protected routes (ต้องมี JWT) ===
	protected := r.Group("/")
	protected.Use(JWTMiddleware())
	{
		// Users (protected)
		protectedUsers := protected.Group("/users")
		protectedUsers.GET("", userController.GetUsers)       // list all users
		protectedUsers.GET("/:id", userController.GetUser)    // get user by id

		protectedPoses := protected.Group("/poses")
		protectedPoses.GET("", poseController.GetPoses)
		protectedPoses.POST("", poseController.CreatePose)

		// Game (protected)
		protectedGame := protected.Group("/game")
		protectedGame.GET("/plan", gameController.GetCurrentGamePlan)
		protectedGame.POST("/complete", gameController.CompleteDay)
	}

	return r
}
