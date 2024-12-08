package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"` // Храним хеш пароля
	CreatedAt time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
}
