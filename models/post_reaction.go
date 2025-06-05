package models

import "time"

type PostReaction struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	PostID    string    `gorm:"type:varchar(36);index"`
	UserID    string    `gorm:"type:varchar(36);index"`
	IsLike    bool
	CreatedAt time.Time
}

func (PostReaction) TableName() string { return "post_reactions" }
