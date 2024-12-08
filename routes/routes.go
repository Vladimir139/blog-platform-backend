package routes

import (
	"blog-platform-backend/controllers"
	"blog-platform-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Публичные роуты
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the blog platform API"})
	})

	// Маршруты для пользователей
	r.GET("/users/:id", controllers.GetUserByID)

	// Группа роутов под авторизацией
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	auth.POST("/posts", controllers.CreatePost)
	auth.PUT("/posts/:id", controllers.UpdatePost)
	auth.DELETE("/posts/:id", controllers.DeletePost)
	auth.POST("/posts/:id/like", controllers.LikePost)
	auth.GET("/user/posts", controllers.GetUserPosts)

	return r
}
