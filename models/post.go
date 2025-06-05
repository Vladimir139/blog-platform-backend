package models

import (
	"time"
)

type Post struct {
	ID           string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title        string    `json:"title"`
	Featured     bool      `json:"featured"`
	Likes    int `json:"likes"`
    Dislikes int `json:"dislikes"`
	UserID       string    `gorm:"type:varchar(36);index" json:"userId"`
	Author       User      `json:"author" gorm:"foreignKey:UserID;references:ID"`
	ShortDesc    string    `json:"shortDesc"`
	Content      string    `json:"content"`
	PreviewImage string    `json:"previewImage"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
}

func (post Post) TableName() string {
	return "posts"
}
