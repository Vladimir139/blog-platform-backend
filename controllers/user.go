package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var user models.User

	if err := database.DB.Preload("Posts").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetMe - возвращает текущего авторизованного пользователя
func GetMe(c *gin.Context) {
	// Получаем user_id из контекста, установленного в мидлваре JWTMiddleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Получаем пользователя из базы данных
	var user models.User
	if err := database.DB.Preload("Posts").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
