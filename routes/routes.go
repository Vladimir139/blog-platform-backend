package routes

import (
	"blog-platform-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Публичные роуты
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)

	// Группа роутов под авторизацией
	auth := r.Group("/")
	// Здесь должен быть миддлвар для проверки JWT токена (например, auth.Use(JWTMiddleware()))
	auth.POST("/posts", controllers.CreatePost)
	auth.PUT("/posts/:id", controllers.UpdatePost)
	auth.DELETE("/posts/:id", controllers.DeletePost)

	return r
}