package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1/auth")
	auth.POST("/register", middleware.Register)
	auth.POST("/login", middleware.Login)
	auth.POST("/logout", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	})

	admin := auth.Group("/users")
	admin.Use(middleware.AuthMiddleware())
	admin = r.Group("/api/v1/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.GET("/users", controllers.GetAllUsers)
		admin.PUT("/:id/ban", controllers.BanUser)
		admin.PUT("/:id/unban", controllers.UnbanUser)
	}
}
