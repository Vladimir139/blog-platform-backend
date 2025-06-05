package models

import "time"

type Comment struct {
	ID        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	Content   string    `json:"content" binding:"required"`
	PostID    string    `gorm:"type:varchar(36);index" json:"postId"`
	UserID    string    `gorm:"type:varchar(36);index" json:"userId"`
	Author    User      `json:"author" gorm:"foreignKey:UserID;references:ID"`
	Likes     int       `json:"likes"`
    Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"createdAt"`
}

func (Comment) TableName() string {
	return "comments"
}
