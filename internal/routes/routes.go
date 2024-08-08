package routes

import (
	"log"
	"time"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/database"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/user"
	"github.com/gin-gonic/gin"
)

// Logger Setup
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

// Router Setup
func SetupRouter() *gin.Engine {

	router := gin.Default()

	// *** MIDDLEWARE ***
	router.Use(LoggerMiddleware())

	db := database.InitDB()

	if db == nil {
		log.Fatalf("DB instance could not be established, DB: %v\n", db)
	}

	// User DI
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	// *** ROUTES ***

	api := router.Group("/api")

	// -- User Routes --
	userRoutes := api.Group("/user")
	userRoutes.POST("/signup", userController.SignUp)
	userRoutes.POST("/signin", userController.SignIn)

	return router
}
