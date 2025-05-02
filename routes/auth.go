package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	auth.POST("/register", middleware.Register)
	auth.POST("/login", middleware.Login)

	admin := auth.Group("/users")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.PUT("/:id/ban", controllers.BanUser)
		admin.PUT("/:id/unban", controllers.UnbanUser)
	}
}
