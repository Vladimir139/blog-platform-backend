package models

import "time"

type User struct {
	ID        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Posts     []Post    `json:"posts" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt"`
}
