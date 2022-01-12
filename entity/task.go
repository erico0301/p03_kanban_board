package entity

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id gorm:primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"null"`
	Status      bool      `json:"status" gorm:"not null"`
	UserID      uint      `json:"user_id"`
	User        *User     `gorm:"foreignKey:UserID"`
	CategoryID  uint      `json:"categoryID"`
	Category    *Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
