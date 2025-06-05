package main

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"blog-platform-backend/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Подключаемся к базе
	database.Connect()

	// Выполняем миграции (создание таблиц в БД)
	err := database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.PostReaction{}, &models.CommentReaction{}, &models.Subscription{}, &models.Notification{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	if err := database.DB.Exec(`
	CREATE UNIQUE INDEX IF NOT EXISTS idx_subs_unique
	  ON user_subscriptions(follower_id, author_id);
	`).Error; err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	// Выполнение сидеров
	// log.Println("Seeding database...")
	// users := seeders.SeedUsers()
	// seeders.SeedPosts(users)

	// Настраиваем роутер
	r := routes.SetupRouter()

	// Запускаем сервер на порту 8080 (по умолчанию)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
