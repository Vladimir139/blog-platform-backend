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
	r.GET("/posts/:id/comments", controllers.GetCommentsByPost)
	r.GET("/posts/popular", controllers.GetPopularPosts)
	r.GET("/users/top", controllers.GetTopAuthors)

	// Публичный маршрут для пользователя
	r.GET("/users/:id", controllers.GetUserByID)

	// Авторы
	r.GET("/authors", controllers.GetAuthors)
	r.GET("/authors/:id", controllers.GetAuthorWithPosts)

	// Маршруты под авторизацией
	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	auth.POST("/posts", controllers.CreatePost)
	auth.PUT("/posts/:id", controllers.UpdatePost)
	auth.DELETE("/posts/:id", controllers.DeletePost)

	auth.POST("/posts/:id/reaction", controllers.ReactPost)
	auth.POST("/comments/:id/reaction", controllers.ReactComment)
	auth.POST("/posts/:id/comments", controllers.CreateComment)

	auth.GET("/users/me/posts", controllers.GetUserPosts)
	auth.GET("/users/me/post-reactions", controllers.GetMyPostReactions)
	auth.GET("/users/me", controllers.GetMe)
	auth.PUT("/users/me", controllers.UpdateMe)
	auth.GET("/users/me/comment-reactions", controllers.GetMyCommentReactions)

	auth.POST("/authors/:id/subscribe", controllers.SubscribeAuthor)
	auth.DELETE("/authors/:id/subscribe", controllers.UnsubscribeAuthor)
	auth.GET("/users/me/subscriptions", controllers.GetMySubscriptions)
	auth.GET("/feed/posts", controllers.GetFeedPosts)

	return r
}
