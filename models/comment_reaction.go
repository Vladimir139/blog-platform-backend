package models

import "time"

type CommentReaction struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	CommentID string    `gorm:"type:varchar(36);index"`
	UserID    string    `gorm:"type:varchar(36);index"`
	IsLike    bool
	CreatedAt time.Time
}

func (CommentReaction) TableName() string { return "comment_reactions" }
