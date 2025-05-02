package routes

import (
	"awesomeProject/controllers"
	_ "awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	auth.POST("/register", middleware.Register)
	auth.POST("/login", middleware.Login)
	auth.PUT("/users/:id/ban", controllers.BanUser)
	auth.PUT("/users/:id/unban", controllers.UnbanUser)
}
