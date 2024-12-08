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
