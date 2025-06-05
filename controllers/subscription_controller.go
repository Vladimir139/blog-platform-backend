package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/* ---------- POST /authors/:id/subscribe ---------- */
func SubscribeAuthor(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	followerID := userID.(string)
	authorID := c.Param("id")

	if followerID == authorID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot subscribe to yourself"})
		return
	}

	sub := models.Subscription{
		ID:         uuid.New().String(),
		FollowerID: followerID,
		AuthorID:   authorID,
		CreatedAt:  time.Now(),
	}
	if err := database.DB.
		Where("follower_id = ? AND author_id = ?", followerID, authorID).
		FirstOrCreate(&sub).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscribed": true})
}

/* ---------- DELETE /authors/:id/subscribe ---------- */
func UnsubscribeAuthor(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	followerID := userID.(string)
	authorID := c.Param("id")

	if err := database.DB.
		Where("follower_id = ? AND author_id = ?", followerID, authorID).
		Delete(&models.Subscription{}).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsubscribe"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"subscribed": false})
}

/* ---------- GET /users/me/subscriptions ---------- */
func GetMySubscriptions(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var ids []string
	if err := database.DB.
		Table("user_subscriptions").
		Select("author_id").
		Where("follower_id = ?", userID).
		Pluck("author_id", &ids).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, ids) // ← массив строк
}

/* ---------- GET /feed/posts?limit=20 ---------- */
func GetFeedPosts(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	limit := 20
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			limit = n
		}
	}

	var posts []models.Post
	err := database.DB.
		Joins(`JOIN user_subscriptions s ON posts.user_id = s.author_id AND s.follower_id = ?`, userID).
		Preload("Author").
		Order("posts.created_at DESC").
		Limit(limit).
		Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feed"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
