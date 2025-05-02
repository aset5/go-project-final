package middleware

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	// 📌 JSON тексеру
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 🔁 Логин қайталанбауы керек
	var existing models.User
	if err := config.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	// 🔐 Құпиясөзді хэштеу
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	// ✅ Жаңа user жасау
	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
		Active:   true, // default: тіркелген кезде актив
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed", "details": err.Error()})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 🚫 Егер user бұғатталған болса
	if !user.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": "Your account is banned"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
