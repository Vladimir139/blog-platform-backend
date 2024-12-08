package models

import (
	"time"
)

// Структура поста
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`      // Заголовок поста
	Featured  bool      `json:"featured"`   // Избранный пост или нет
	Likes     int       `json:"likes"`      // Количество лайков (int достаточно, если не будет огромных значений)
	UserID    uint      `json:"user_id"`    // ID автора (Foreign Key на пользователя)
	ShortDesc string    `json:"short_desc"` // Короткое описание для превью
	Content   string    `json:"content"`    // Сам контент редактора (хранится как строка, например JSON от wysiwyg)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
