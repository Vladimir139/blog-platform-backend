package models

import "time"

type Notification struct {
	ID        string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID    string    `gorm:"type:varchar(36);index"      json:"userId"`
	AuthorID  string    `gorm:"type:varchar(36);index"      json:"authorId"`
	PostID    string    `gorm:"type:varchar(36);index"      json:"postId"`
	Message   string    `json:"message"`
	IsRead    bool      `gorm:"default:false"               json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`

	Author User `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	Post   Post `json:"post"   gorm:"foreignKey:PostID;references:ID"`
}

func (Notification) TableName() string { return "notifications" }
