package models

import (
	"time"
)

type UserToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	Provider  string    `gorm:"type:varchar(50);not null" json:"provider"` // google, etc.
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
