package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`                              // хеш пароля
	Posts     []Post    `json:"posts" gorm:"foreignKey:UserID"` // Связь "один ко многим"
	CreatedAt time.Time `json:"createdAt"`
}
