package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"constraint:OnDelete:CASCADE" json:"author"`
	PostID    uint      `json:"post_id"`
	Post      Post      `gorm:"constraint:OnDelete:CASCADE" json:"post"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
