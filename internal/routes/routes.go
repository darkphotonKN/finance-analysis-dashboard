package routes

import (
	"log"
	"time"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/controllers"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// *** MIDDLEWARE ***
	r.Use(LoggerMiddleware())

	// *** ROUTES ***

	// -- User Routes --
	r.POST("/api/signup", controllers.SignUp)

	return r
}
