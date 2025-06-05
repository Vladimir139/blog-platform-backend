package models

import "time"

// Subscription — связь «follower → author».
type Subscription struct {
	ID         string `gorm:"type:varchar(36);primaryKey"`
	FollowerID string `gorm:"type:varchar(36);index"` // кто подписан
	AuthorID   string `gorm:"type:varchar(36);index"` // на кого
	CreatedAt  time.Time

	Follower User `gorm:"foreignKey:FollowerID;references:ID"`
	Author   User `gorm:"foreignKey:AuthorID;references:ID"`
}

// уникальный индекс задаём тегом или миграцией
func (Subscription) TableName() string { return "user_subscriptions" }
