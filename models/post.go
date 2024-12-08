package models

import (
	"time"
)

// Структура поста
type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Featured  bool      `json:"featured"`
	Likes     int       `json:"likes"`
	UserID    uint      `json:"userId"`
	Author    User      `json:"author" gorm:"foreignKey:UserID"`
	ShortDesc string    `json:"shortDesc"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
