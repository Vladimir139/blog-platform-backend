package seeders

import (
	"blog-platform-backend/database"
	"blog-platform-backend/models"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SeedUsers создает тестовых пользователей
func SeedUsers() []models.User {
	users := []models.User{
		{
			ID:        uuid.New().String(),
			FirstName: "Иван",
			LastName:  "Иванов",
			Email:     "ivan@example.com",
			Password:  hashPassword("password123"),
		},
		{
			ID:        uuid.New().String(),
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
			ID:           uuid.New().String(),
			Title:        "Уроки лидерства от Билла Уолша",
			Featured:     true,
			ShortDesc:    "Узнайте секреты превращения команды с 2-14 в 3-кратных победителей Супербоула.",
			Content:      "<p>Узнайте секреты превращения команды с 2-14 в 3-кратных победителей Супербоула.</p>",
			PreviewImage: "/images/posts/1.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Искусство войны в бизнесе",
			Featured:     false,
			ShortDesc:    "Применение стратегий Сунь Цзы к современным бизнес-задачам.",
			Content:      "<p>Применение стратегий Сунь Цзы к современным бизнес-задачам.</p>",
			PreviewImage: "/images/posts/2.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 2, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Осознанность для продуктивности",
			Featured:     false,
			ShortDesc:    "Как практики осознанности повышают эффективность работы.",
			Content:      "<p>Как практики осознанности повышают эффективность работы.</p>",
			PreviewImage: "/images/posts/3.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 3, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Исследуя космос",
			Featured:     false,
			ShortDesc:    "Путешествие по последним открытиям в космических исследованиях.",
			Content:      "<p>Путешествие по последним открытиям в космических исследованиях.</p>",
			PreviewImage: "/images/posts/4.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 4, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Здоровое питание на бюджете",
			Featured:     false,
			ShortDesc:    "Советы и хитрости для питательных блюд без лишних затрат.",
			Content:      "<p>Советы и хитрости для питательных блюд без лишних затрат.</p>",
			PreviewImage: "/images/posts/5.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Будущее искусственного интеллекта",
			Featured:     false,
			ShortDesc:    "Прогнозы и возможности ИИ в нашей жизни.",
			Content:      "<p>Прогнозы и возможности ИИ в нашей жизни.</p>",
			PreviewImage: "/images/posts/6.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 6, 18, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Путешествие по миру виртуально",
			Featured:     false,
			ShortDesc:    "Исследуйте мировые достопримечательности, не выходя из дома.",
			Content:      "<p>Исследуйте мировые достопримечательности, не выходя из дома.</p>",
			PreviewImage: "/images/posts/7.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2023, 7, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Основы устойчивого образа жизни",
			Featured:     false,
			ShortDesc:    "Простые шаги к снижению экологического следа.",
			Content:      "<p>Простые шаги к снижению экологического следа.</p>",
			PreviewImage: "/images/posts/8.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Мастерство кулинарии",
			Featured:     false,
			ShortDesc:    "От новичка до шефа: ваш путь к успеху в кулинарии.",
			Content:      "<p>От новичка до шефа: ваш путь к успеху в кулинарии.</p>",
			PreviewImage: "/images/posts/9.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Технологии завтрашнего дня",
			Featured:     true,
			ShortDesc:    "Как современные технологии меняют будущее.",
			Content:      "<p>Погружение в мир современных технологий и их влияние на нашу жизнь.</p>",
			PreviewImage: "/images/posts/10.jpg",
			UserID:       users[1].ID,
			CreatedAt:    time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Искусство фотографии",
			Featured:     false,
			ShortDesc:    "Узнайте секреты создания захватывающих снимков.",
			Content:      "<p>Советы и техники для улучшения ваших навыков фотографии.</p>",
			PreviewImage: "/images/posts/1.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2023, 11, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Психология успеха",
			Featured:     true,
			ShortDesc:    "Как мышление влияет на достижение целей.",
			Content:      "<p>Понимание того, как мышление и привычки помогают достигать больших высот.</p>",
			PreviewImage: "/images/posts/2.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:           uuid.New().String(),
			Title:        "Сила привычек",
			Featured:     false,
			ShortDesc:    "Как маленькие изменения приводят к большим результатам.",
			Content:      "<p>Пошаговый разбор того, как формировать и поддерживать полезные привычки.</p>",
			PreviewImage: "/images/posts/3.jpg",
			UserID:       users[0].ID,
			CreatedAt:    time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
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
