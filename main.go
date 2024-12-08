package main

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"blog-platform-backend/routes"
	"blog-platform-backend/seeders"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Подключаемся к базе
	database.Connect()

	// Выполняем миграции (создание таблиц в БД)
	err := database.DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// Выполнение сидеров
	log.Println("Seeding database...")
	users := seeders.SeedUsers()
	seeders.SeedPosts(users)

	// Настраиваем роутер
	r := routes.SetupRouter()

	// Запускаем сервер на порту 8080 (по умолчанию)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
