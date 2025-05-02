package main

import (
	"awesomeProject/config"
	_ "awesomeProject/middleware"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	routes.RegisterProductRoutes(r)
	routes.RegisterAuthRoutes(r)

	r.Run(":8080")
}
