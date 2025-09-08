package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID          string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID      string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Amount      float64        `gorm:"not null" json:"amount"`
	Category    string         `gorm:"type:varchar(50)" json:"category"`
	Description string         `gorm:"type:text" json:"description"`
	Date        time.Time      `gorm:"not null" json:"date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
