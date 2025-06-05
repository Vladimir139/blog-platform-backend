package controllers

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ----------  POST REACTION ----------

type reactionInput struct {
	IsLike bool `json:"isLike"` // true = like, false = dislike
}

func ReactPost(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDStr := userID.(string)
	postID := c.Param("id")

	var body reactionInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var post models.Post
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&post, "id = ?", postID).Error; err != nil {
			return err
		}

		var react models.PostReaction
		res := tx.Where("post_id = ? AND user_id = ?", postID, userIDStr).First(&react)

		switch {
		case res.Error == gorm.ErrRecordNotFound:
			// ещё не реагировал
			react = models.PostReaction{
				ID:        uuid.New().String(),
				PostID:    postID,
				UserID:    userIDStr,
				IsLike:    body.IsLike,
				CreatedAt: time.Now(),
			}
			if err := tx.Create(&react).Error; err != nil {
				return err
			}
			if body.IsLike {
				post.Likes++
			} else {
				post.Dislikes++
			}

		case res.Error == nil:
			if react.IsLike == body.IsLike {
				// повторное нажатие → снимаем реакцию
				if react.IsLike {
					post.Likes--
				} else {
					post.Dislikes--
				}
				if err := tx.Delete(&react).Error; err != nil {
					return err
				}
			} else {
				// меняем like ↔ dislike
				if body.IsLike {
					post.Likes++
					post.Dislikes--
				} else {
					post.Likes--
					post.Dislikes++
				}
				react.IsLike = body.IsLike
				if err := tx.Save(&react).Error; err != nil {
					return err
				}
			}
		default:
			return res.Error
		}
		return tx.Save(&post).Error
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to react"})
		return
	}

	// Возвращаем свежие счётчики
	var post models.Post
	if err := database.DB.
		Preload("Author").
		Preload("Comments").
		Preload("Comments.Author").
		First(&post, "id = ?", postID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post not found after update"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// ----------  COMMENT REACTION ----------

// controllers/reaction_controller.go

// ReactComment — поставить / снять like-/dislike-реакцию на комментарий
//
//	POST /comments/:id/reaction   { "isLike": true | false }
func ReactComment(c *gin.Context) {
	/* ----------  0. Авторизация  ---------- */
	uid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := uid.(string)
	commentID := c.Param("id")

	/* ----------  1. Парсим тело запроса  ---------- */
	var body reactionInput // { isLike: bool }
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/* ----------  2. Транзакция «реакция + счётчики»  ---------- */
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// a) блокируем комментарий на время изменения
		var comment models.Comment
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&comment, "id = ?", commentID).Error; err != nil {
			return err
		}

		// b) ищем текущую реакцию пользователя
		var react models.CommentReaction
		q := tx.Where("comment_id = ? AND user_id = ?", commentID, userID).
			First(&react)

		switch {
		/* --- пользователь ещё не реагировал --- */
		case q.Error == gorm.ErrRecordNotFound:
			react = models.CommentReaction{
				ID:        uuid.New().String(),
				CommentID: commentID,
				UserID:    userID,
				IsLike:    body.IsLike,
				CreatedAt: time.Now(),
			}
			if err := tx.Create(&react).Error; err != nil {
				return err
			}
			if body.IsLike {
				comment.Likes++
			} else {
				comment.Dislikes++
			}

		/* --- реакция уже есть --- */
		case q.Error == nil:
			/* 1) повторный клик по той же кнопке → отменяем реакцию */
			if react.IsLike == body.IsLike {
				if react.IsLike {
					comment.Likes--
				} else {
					comment.Dislikes--
				}
				if err := tx.Delete(&react).Error; err != nil {
					return err
				}

				/* 2) переключение like ↔ dislike */
			} else {
				if body.IsLike {
					comment.Likes++
					comment.Dislikes--
				} else {
					comment.Likes--
					comment.Dislikes++
				}
				react.IsLike = body.IsLike
				if err := tx.Save(&react).Error; err != nil {
					return err
				}
			}

		/* --- любая другая ошибка выборки --- */
		default:
			return q.Error
		}

		// c) сохраняем обновлённые счётчики
		return tx.Save(&comment).Error
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to react"})
		return
	}

	/* ----------  3. Отдаём «полный» комментарий  ---------- */
	var out models.Comment
	if err := database.DB.
		Preload("Author", func(db *gorm.DB) *gorm.DB { // можно сузить поля автора
			return db.Select("id, first_name, last_name")
		}).
		First(&out, "id = ?", commentID).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, out)
}

func GetMyPostReactions(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var reactions []models.PostReaction
	if err := database.DB.Where("user_id = ?", userID).Find(&reactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reactions"})
		return
	}

	type Response struct {
		PostID string `json:"postId"`
		Type   string `json:"type"` // "like" | "dislike"
	}

	var out []Response
	for _, r := range reactions {
		t := "dislike"
		if r.IsLike {
			t = "like"
		}
		out = append(out, Response{PostID: r.PostID, Type: t})
	}

	c.JSON(http.StatusOK, out)
}

func GetMyCommentReactions(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var reactions []models.CommentReaction
	if err := database.DB.Where("user_id = ?", userID).
		Find(&reactions).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reactions"})
		return
	}

	type resp struct {
		CommentID string `json:"commentId"`
		Type      string `json:"type"` // "like" | "dislike"
	}

	out := make([]resp, 0, len(reactions))
	for _, r := range reactions {
		t := "dislike"
		if r.IsLike {
			t = "like"
		}
		out = append(out, resp{CommentID: r.CommentID, Type: t})
	}

	c.JSON(http.StatusOK, out)
}
