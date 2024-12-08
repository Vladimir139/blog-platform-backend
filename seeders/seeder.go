package seeders

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedUsers создает тестовых пользователей
func SeedUsers() []models.User {
	users := []models.User{
		{
			FirstName: "Иван",
			LastName:  "Иванов",
			Email:     "ivan@example.com",
			Password:  hashPassword("password123"),
		},
		{
			FirstName: "Анна",
			LastName:  "Петрова",
			Email:     "anna@example.com",
			Password:  hashPassword("password456"),
		},
	}

	for i := range users {
		if err := database.DB.Create(&users[i]).Error; err != nil {
			log.Fatalf("Failed to seed user: %v", err)
		}
		fmt.Printf("User seeded: %s\n", users[i].Email)
	}

	return users
}

// SeedPosts создает тестовые посты
func SeedPosts(users []models.User) {
	posts := []models.Post{
		{
			Title:     "Первый пост",
			Featured:  true,
			ShortDesc: "Это первый тестовый пост",
			Content:   "<p>Полный текст первого поста...</p>",
			UserID:    users[0].ID, // Используем ID первого пользователя
		},
		{
			Title:     "Второй пост",
			Featured:  false,
			ShortDesc: "Это второй тестовый пост",
			Content:   "<p>Полный текст второго поста...</p>",
			UserID:    users[1].ID, // Используем ID второго пользователя
		},
	}

	for _, post := range posts {
		if err := database.DB.Create(&post).Error; err != nil {
			log.Fatalf("Failed to seed post: %v", err)
		}
		fmt.Printf("Post seeded: %s\n", post.Title)
	}
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hash)
}
