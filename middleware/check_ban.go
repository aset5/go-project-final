package middleware

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckIfBanned() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDAny, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		var userID uint
		switch v := userIDAny.(type) {
		case float64:
			userID = uint(v)
		case uint:
			userID = v
		case int:
			userID = uint(v)
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
			return
		}

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if !user.Active {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Your account is banned"})
			return
		}

		c.Next()
	}
}
