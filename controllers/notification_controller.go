package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* GET /users/me/notifications?unread=true&limit=20&offset=0 */
func GetMyNotifications(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	limit, offset := 20, 0
	if v, _ := strconv.Atoi(c.DefaultQuery("limit", "20")); v > 0 {
		limit = v
	}
	if v, _ := strconv.Atoi(c.DefaultQuery("offset", "0")); v >= 0 {
		offset = v
	}

	query := database.DB.
		Where("user_id = ?", userID).
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, first_name, last_name")
		}).
		Preload("Post")

	if c.Query("unread") == "true" {
		query = query.Where("is_read = false")
	}

	var list []models.Notification
	if err := query.
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&list).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, list)
}

/* PUT /notifications/:id/read */
func MarkNotificationRead(c *gin.Context) {
	userID, _ := c.Get("user_id")
	notifID := c.Param("id")

	if err := database.DB.
		Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", notifID, userID).
		Update("is_read", true).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	c.Status(http.StatusOK)
}

/* PUT /notifications/read-all */
func MarkAllRead(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if err := database.DB.
		Model(&models.Notification{}).
		Where("user_id = ?", userID).
		Update("is_read", true).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}
	c.Status(http.StatusOK)
}
