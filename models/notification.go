package models

import "time"

type Notification struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	UserID    string `gorm:"type:varchar(36);index"` // получатель
	AuthorID  string `gorm:"type:varchar(36);index"` // кто создал пост
	PostID    string `gorm:"type:varchar(36);index"`
	Message   string
	IsRead    bool `gorm:"default:false"`
	CreatedAt time.Time

	Author User `gorm:"foreignKey:AuthorID;references:ID"`
	Post   Post `gorm:"foreignKey:PostID;references:ID"`
}

func (Notification) TableName() string { return "notifications" }
