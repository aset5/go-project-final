package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/products", controllers.CreateProduct)
		protected.PUT("/products/:id", controllers.UpdateProduct)
		protected.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
