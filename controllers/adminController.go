package controllers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BanUser(c *gin.Context) {
	roleAny, exists := c.Get("userRole")
	if !exists || roleAny != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can ban users"})
		return
	}

	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.IsBanned = true
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to ban user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User banned"})
}

func UnbanUser(c *gin.Context) {
	roleAny, exists := c.Get("userRole")
	if !exists || roleAny != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can unban users"})
		return
	}

	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.IsBanned = false
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unban user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User unbanned"})
}
