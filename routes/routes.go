package routes

import (
	"debt_tracker/controllers"
	"debt_tracker/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/me", middleware.AuthMiddleware(), func(c *gin.Context) {
			userID := c.MustGet("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}
}
