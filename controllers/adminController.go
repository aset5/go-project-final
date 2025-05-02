package controllers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BanUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Active = false
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User banned"})
}

func UnbanUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Active = true
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User unbanned"})
}
