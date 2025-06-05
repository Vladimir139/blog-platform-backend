package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /posts/popular?limit=10
func GetPopularPosts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var posts []models.Post

	database.DB.
		Preload("Author").
		Order("likes - dislikes DESC, created_at DESC").
		Limit(limit).
		Find(&posts)

	c.JSON(http.StatusOK, posts)
}

// GET /users/top?limit=10
func GetTopAuthors(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	type authorLikes struct {
		models.User
		Total int64 `json:"totalLikes"`
	}

	var res []authorLikes
	database.DB.Table("users").
		Select("users.*, COALESCE(SUM(posts.likes - posts.dislikes),0) AS total").
		Joins("LEFT JOIN posts ON posts.user_id = users.id").
		Group("users.id").
		Order("total DESC").
		Limit(limit).
		Scan(&res)

	c.JSON(http.StatusOK, res)
}
