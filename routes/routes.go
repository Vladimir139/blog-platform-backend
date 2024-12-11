package routes

import (
	"blog-platform-backend/controllers"
	"blog-platform-backend/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Добавляем CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Разрешаем запросы с этого источника
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Публичные роуты
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the blog platform API"})
	})

	// Эндпоинты аутентификации
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/refresh", controllers.Refresh)

	// Публичные маршруты для постов
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)

	// Публичный маршрут для пользователя
	r.GET("/users/:id", controllers.GetUserByID)

	// Маршруты под авторизацией
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	auth.POST("/posts", controllers.CreatePost)
	auth.PUT("/posts/:id", controllers.UpdatePost)
	auth.DELETE("/posts/:id", controllers.DeletePost)
	auth.POST("/posts/:id/like", controllers.LikePost)
	auth.GET("/users/me/posts", controllers.GetUserPosts)
	auth.GET("/users/me", controllers.GetMe)
	auth.PUT("/users/me", controllers.UpdateMe)

	return r
}
