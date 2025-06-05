package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePost - создание поста
func CreatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		Title        string `json:"title" binding:"required"`
		Featured     bool   `json:"featured"`
		ShortDesc    string `json:"shortDesc"`
		Content      string `json:"content" binding:"required"`
		PreviewImage string `json:"previewImage"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:        input.Title,
		Featured:     input.Featured,
		ShortDesc:    input.ShortDesc,
		Content:      input.Content,
		PreviewImage: input.PreviewImage,
		UserID:       userID.(string), // UserID как строка
		ID:           uuid.New().String(),
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

	if err := database.DB.Preload("Author").Preload("Comments").Preload("Comments.Author").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostByID - получить пост по ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Preload("Author").Preload("Comments").Preload("Comments.Author").First(&post, "id = ?", id).Error; err != nil {
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
	userIDStr := userID.(string)

	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userIDStr {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	var input struct {
		Title        *string `json:"title"`
		Featured     *bool   `json:"featured"`
		ShortDesc    *string `json:"shortDesc"`
		Content      *string `json:"content"`
		PreviewImage *string `json:"previewImage"`
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
	if input.PreviewImage != nil {
		post.PreviewImage = *input.PreviewImage
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
	userIDStr := userID.(string)

	id := c.Param("id")
	var post models.Post
	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userIDStr {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func LikePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.DB.First(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post.Likes += 1
	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"likes": post.Likes})
}

func GetUserPosts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDStr := userID.(string)

	var posts []models.Post
	if err := database.DB.Where("user_id = ?", userIDStr).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
