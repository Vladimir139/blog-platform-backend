package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* ------------ GET /authors -------------- */
// ?limit=20&offset=0
func GetAuthors(c *gin.Context) {
	limit := 20
	offset := 0

	if v := c.Query("limit"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			limit = n
		}
	}
	if v := c.Query("offset"); v != "" {
		if n, _ := strconv.Atoi(v); n >= 0 {
			offset = n
		}
	}

	var authors []models.User
	err := database.DB.
		Select("id, first_name, last_name, created_at").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&authors).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
		return
	}
	c.JSON(http.StatusOK, authors)
}

/* -------- GET /authors/:id (+ его посты) -------- */
// ?limit=20&offset=0
func GetAuthorWithPosts(c *gin.Context) {
	authorID := c.Param("id")

	var author models.User
	if err := database.DB.
		Select("id, first_name, last_name, created_at").
		First(&author, "id = ?", authorID).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	limit := 20
	offset := 0
	if v := c.Query("limit"); v != "" {
		if n, _ := strconv.Atoi(v); n > 0 {
			limit = n
		}
	}
	if v := c.Query("offset"); v != "" {
		if n, _ := strconv.Atoi(v); n >= 0 {
			offset = n
		}
	}

	var posts []models.Post
	err := database.DB.
		Where("user_id = ?", authorID).
		Preload("Author").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"author": author,
		"posts":  posts,
	})
}
