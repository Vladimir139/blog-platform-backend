package controllers

import (
	"net/http"
	"strconv"

	"blog-platform-backend/database"
	"blog-platform-backend/models"

	"github.com/gin-gonic/gin"
)

// CreatePost - создание поста
func CreatePost(c *gin.Context) {
	// Предполагается, что пользователь уже аутентифицирован
	// Предположим, что в контексте (через мидлвару JWT) есть user_id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Title     string `json:"title" binding:"required"`
		Featured  bool   `json:"featured"`
		ShortDesc string `json:"short_desc"`
		Content   string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:     input.Title,
		Featured:  input.Featured,
		ShortDesc: input.ShortDesc,
		Content:   input.Content,
		UserID:    userID.(uint), // приводим к uint
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPosts - получить все посты
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPostByID - получить пост по ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost - обновление поста
func UpdatePost(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Разрешить обновление только, если текущий пользователь - автор поста
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	var input struct {
		Title     *string `json:"title"`
		Featured  *bool   `json:"featured"`
		ShortDesc *string `json:"short_desc"`
		Content   *string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Featured != nil {
		post.Featured = *input.Featured
	}
	if input.ShortDesc != nil {
		post.ShortDesc = *input.ShortDesc
	}
	if input.Content != nil {
		post.Content = *input.Content
	}

	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost - удалить пост
func DeletePost(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr) // Допустим, нет ошибок
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
