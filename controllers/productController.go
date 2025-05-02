package controllers

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if result := config.DB.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}
	product.UserID = userID.(uint)

	if product.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if config.DB.First(&product, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = input.Name
	product.Price = input.Price
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if result := config.DB.Delete(&models.Product{}, id); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
